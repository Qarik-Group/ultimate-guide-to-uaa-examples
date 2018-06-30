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
20
```

The airports app attempts to perform a "whoami" request with the UAA to confirm tht the access token is valid, and dumps the `/userinfo` output to its logs:

```json
{"user_id":"871a53d4-11ac-437b-ba28-c84b795b0221","user_name":"drnic","name":"Dr Nic Williams","given_name":"Dr Nic","family_name":"Williams","email":"drnic@starkandwayne.com","email_verified":true,"previous_logon_time":1530334456892,"sub":"871a53d4-11ac-437b-ba28-c84b795b0221"}
```