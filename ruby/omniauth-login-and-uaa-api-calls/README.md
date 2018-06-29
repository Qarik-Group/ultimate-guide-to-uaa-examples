# Ruby OmniAuth - Simple Login Only

First, create a UAA client:

```text
uaa-deployment auth-client
uaa create-client omniauth-login-and-uaa-api-calls -s omniauth-login-and-uaa-api-calls \
  --authorized_grant_types authorization_code,refresh_token \
  --scope openid \
  --redirect_uri http://localhost:9292/auth/cloudfoundry/callback
```

Next, setup `$UAA_URL`/`$UAA_CA_CERT`:

```text
source <(path/to/uaa-deployment/bin/uaa-deployment env)
```

Next, run the app:

```text
bundle
bundle exec rackup
```

Visit https://localhost:9292 and commence the login sequence to be redirected to the UAA, and ultimately returned to the app.

