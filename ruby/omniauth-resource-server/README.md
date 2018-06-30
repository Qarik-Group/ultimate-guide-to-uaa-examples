# Scope access to API

First, add `uaa-deployment` into the `$PATH` and setup `$UAA_URL`/`$UAA_CA_CERT`:

```text
source <(path/to/uaa-deployment/bin/uaa-deployment env)
uaa-deployment auth-client
```

Create some demo users:

```text
uaa create-user airports-all --password airports-all \
  --email airports-all@example.com \
  --givenName "Airports" --familyName "All"
uaa create-user airports-50 --password airports-50 \
  --email airports-50@example.com \
  --givenName "Airports" --familyName "50"
uaa create-user airports-no-scope --password airports-no-scope \
  --email airports-no-scope@example.com \
  --givenName "Airports" --familyName "No Scope"
```

Run the Airports API application:

```text
bundle
bundle exec rackup
```

You can access the API as a non-authenticated guest and will receive a JSON array of 10 airports:

```text
$ curl -s http://localhost:9292 | jq length
10
```

Now create a dedicated Airports client:

```text
uaa create-client airports -s airports \
  --authorized_grant_types password,authorization_code,refresh_token \
  --scope airports.all,airports.50,openid \
  --redirect_uri http://localhost:9876
```

Next, authenticate one of the users with the UAA to get an "access_token":

```text
uaa get-password-token airports -s airports -u airports-no-scope -p airports-no-scope
uaa context --access_token
```

Pass the access_token into the `-H 'Authorization: bearer <access_token>'` below:

```text
$ curl -s -H "Authorization: bearer $(uaa context --access_token)" http://localhost:9292 | jq length
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
uaa-deployment auth-client
uaa create-group airports.all -d "Display all airports"
uaa create-group airports.50 -d "Display 50 airports"
```

Grant your `airports-all` user access to all airports (implemented via scope `airports-all`), and `airports-50` user access to scope `airports.50`.

```text
uaa add-member airports.all airports-all
uaa add-member airports.50 airports-50
```

Login as `airports-50` user and see that the Airports API now returns 50 results:

```text
uaa get-password-token airports -s airports -u airports-50 -p airports-50
curl -s -H "Authorization: bearer $(uaa context --access_token)" http://localhost:9292 | jq length
```

Finally, login as `airports-all` user and see that the Airports API now returns all 297 results:

```text
uaa get-password-token airports -s airports -u airports-all -p airports-all
curl -s -H "Authorization: bearer $(uaa context --access_token)" http://localhost:9292 | jq length
```
