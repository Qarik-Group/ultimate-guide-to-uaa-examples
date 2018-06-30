# Applications filter results from UAA user scopes

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
uaa context --auth_header
```

Pass the `uaa context --auth_header` into the `-H 'Authorization:'` below:

```text
$ curl -s -H "Authorization: $(uaa context --auth_header)" http://localhost:9292 | jq length
30
```

The airports app attempts to decode the access token, to confirm it originated from its UAA and used the `airports` client ID, and if successful will expand the limit to 20 airports. The decoded token is dumped to the logs:

```json
{"jti":"ce148a7201634537b0f9da4af24ba91e","sub":"2b4f7895-9a67-4274-a6fd-d2257d492e00","scope":["openid"],"client_id":"airports","cid":"airports","azp":"airports","grant_type":"password","user_id":"2b4f7895-9a67-4274-a6fd-d2257d492e00","origin":"uaa","user_name":"airports-no-scope","email":"airports-no-scope@example.com","auth_time":1530388796,"rev_sig":"d5e8bdec","iat":1530388796,"exp":1530431996,"iss":"https://192.168.50.6:8443/oauth/token","zid":"uaa","aud":["openid","airports"]}
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
curl -s -H "Authorization: $(uaa context --auth_header)" http://localhost:9292 | jq length
```

Finally, login as `airports-all` user and see that the Airports API now returns all 297 results:

```text
uaa get-password-token airports -s airports -u airports-all -p airports-all
curl -s -H "Authorization: $(uaa context --auth_header)" http://localhost:9292 | jq length
```

## Docker

```text
docker build -t starkandwayne/omniauth-resource-server .
docker run -ti -p 9292:9292 -e UAA_URL=$UAA_URL -e UAA_CA_CERT=$UAA_CA_CERT starkandwayne/omniauth-resource-server
```