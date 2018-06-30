# Scope access to API

```text
bundle
bundle exec rackup
```

If you access the API as a guest you will receive a JSON array of 10 airports:

```text
$ curl -s http://localhost:9292 | jq "length"
10
```

Next, authenticate with the UAA to get an "access_token":

```text
uaa get-authcode-token uaa-cli-authcode -s uaa-cli-authcode --port 9876 -v
```

Pass the access_token into the `-H 'Authorization: bearer <access_token>'` below:

```text
$ curl -s http://localhost:9292 -H 'Authorization: bearer <access_token>' | jq "length"
30
```

The airports app attempts to decode the access token and if successful will expand the limit to 20 airports. The decoded token is dumped to the logs:

```json
{"jti":"c695b5fca6354ba0a1e0175894056405","sub":"871a53d4-11ac-437b-ba28-c84b795b0221","scope":["openid"],"client_id":"uaa-cli-authcode","cid":"uaa-cli-authcode","azp":"uaa-cli-authcode","grant_type":"authorization_code","user_id":"871a53d4-11ac-437b-ba28-c84b795b0221","origin":"uaa","user_name":"drnic","email":"drnic@starkandwayne.com","auth_time":1530340069,"rev_sig":"1dc7b713","iat":1530340077,"exp":1530383277,"iss":"https://192.168.50.6:8443/oauth/token","zid":"uaa","aud":["uaa-cli-authcode","openid"]}
```

The airports app attempts to perform a "whoami" request with the UAA, and if successful increases the API limit to 20. It also dumps the `/userinfo` output to its logs so we can see who is accessing our API:

```json
{"user_id":"871a53d4-11ac-437b-ba28-c84b795b0221","user_name":"drnic","name":"Dr Nic Williams","given_name":"Dr Nic","family_name":"Williams","email":"drnic@starkandwayne.com","email_verified":true,"previous_logon_time":1530334456892,"sub":"871a53d4-11ac-437b-ba28-c84b795b0221"}
```

Create some user groups/client scopes that might grant to users:

```text
uaa create-group airports.all -d "Display all airports"
uaa create-group airports.50 -d "Display 50 airports"
```

Now create a dedicated Airports client:

```text
uaa create-client airports -s airports \
  --authorized_grant_types authorization_code,refresh_token \
  --scope airports.all,airports.50,openid \
  --redirect_uri http://localhost:9876
```

Grant your user access to all airports:

```text
uaa add-member airports.all drnic
```

Login again:

```text
uaa get-authcode-token airports -s airports --port 9876 -v
```

Copy the new access token from the output into your curl command:

```text
$ curl -s http://localhost:9292 -H 'Authorization: bearer <new access_token>' | jq "length"
297
```

