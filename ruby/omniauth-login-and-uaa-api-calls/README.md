# Ruby OmniAuth - Simple Login Only

First, add `uaa-deployment` into the `$PATH` and setup `$UAA_URL`/`$UAA_CA_CERT`:

```text
source <(path/to/uaa-deployment/bin/uaa-deployment env)
```

Next, create a UAA client that includes both `openid` and `scim.read` scopes:

```text
uaa-deployment auth-client
uaa create-client omniauth-login-and-uaa-api-calls -s omniauth-login-and-uaa-api-calls \
  --authorized_grant_types authorization_code,refresh_token \
  --scope openid,scim.read \
  --redirect_uri http://localhost:9292/auth/cloudfoundry/callback
```

Add user `drnic` to group `scim.read`:

```text
uaa add-member scim.read drnic
```

If `drnic` not in group `scim.read` then home page will show "drnic@starkandwayne.com ( openid )" instead of "drnic@starkandwayne.com ( openid scim.read )", and "My UAA Info" link will produce an error relating to required scopes.


Next, run the app:

```text
bundle
bundle exec rackup
```

Visit https://localhost:9292 and commence the login sequence to be redirected to the UAA, and ultimately returned to the app.

