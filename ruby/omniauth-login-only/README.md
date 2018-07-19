# Ruby OmniAuth - Simple Login Only

First, create a UAA client:

```text
source <(path/to/uaa-deployment-cf/bin/u env)
u auth-client
uaa create-client omniauth-login-only -s omniauth-login-only \
  --authorized_grant_types authorization_code,refresh_token \
  --scope openid \
  --redirect_uri http://localhost:9292/auth/cloudfoundry/callback,http://127.0.0.1:9292/auth/cloudfoundry/callback
```

Next, setup `$UAA_URL`:

```text
source <(path/to/uaa-deployment-cf/bin/u env)
```

Next, run the app:

```text
bundle
bundle exec rackup
```

Visit https://localhost:9292 and commence the login sequence to be redirected to the UAA, and ultimately returned to the app.

It will output the data returned from the UAA response token, for example:

```json
{"provider":"cloudfoundry","uid":"10661d47-6d08-4872-81aa-7868327cff0b","info":{"name":"Dr Nic Williams","email":"drnic@starkandwayne.com","first_name":"Dr Nic","last_name":"Williams"},"credentials":{"token":"eyJhbGciOiJSUzI1NiIsImtpZCI6ImxlZ2FjeS10b2tlbi1rZXkiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiIwMDcyMDY1NThjYjM0ODJmYjExMDk4NjM3OTQ0MTQ3ZCIsIm5vbmNlIjoiODdlY2FhMGRmNTRhZmE2YmVlMTRmYzQyYWE0NDBkNTciLCJzdWIiOiIxMDY2MWQ0Ny02ZDA4LTQ4NzItODFhYS03ODY4MzI3Y2ZmMGIiLCJzY29wZSI6WyJvcGVuaWQiXSwiY2xpZW50X2lkIjoib21uaWF1dGgtZXhhbXBsZSIsImNpZCI6Im9tbmlhdXRoLWV4YW1wbGUiLCJhenAiOiJvbW5pYXV0aC1leGFtcGxlIiwiZ3JhbnRfdHlwZSI6ImF1dGhvcml6YXRpb25fY29kZSIsInVzZXJfaWQiOiIxMDY2MWQ0Ny02ZDA4LTQ4NzItODFhYS03ODY4MzI3Y2ZmMGIiLCJvcmlnaW4iOiJ1YWEiLCJ1c2VyX25hbWUiOiJkcm5pYyIsImVtYWlsIjoiZHJuaWNAc3RhcmthbmR3YXluZSIsImF1dGhfdGltZSI6MTUzMDI3MDI1MiwicmV2X3NpZyI6Ijg0MmQ5N2RkIiwiaWF0IjoxNTMwMjcxMDM2LCJleHAiOjE1MzAzMTQyMzYsImlzcyI6Imh0dHBzOi8vMTkyLjE2OC41MC42Ojg0NDMvb2F1dGgvdG9rZW4iLCJ6aWQiOiJ1YWEiLCJhdWQiOlsib3BlbmlkIiwib21uaWF1dGgtZXhhbXBsZSJdfQ.rdkpQmvYXUWKTf28Go24-Vju7rkGLRzgXQfBWV5vTOzMxbNdbwdmhyCciahi1gbFyR-XzaHtboOK0vF2VX8BPd9_6zfB_DoVbQb3SV-QyOFVDbUNOQ68GGpw0ct2KsTxKeIygf_2kEfmZ4nQ6W83taWJU6cFMEOUI-Z3vUzPD09_HCBuTkOj2Nr9EdenrnFZcnu4y5r6Pw8YdE9KGoJnMiPfMYF1ilBzafDg85t_Wy6aplZr5GWpQrJhSUUokQZ1wKUwzEj_YzDV9FWmUzAqMJxEuXDcUkKegCI394qhyHjel20G_nSPrp8nWaSC8ee1RyNtIP1efBPOx5iF8qGGGg","refresh_token":"eyJhbGciOiJSUzI1NiIsImtpZCI6ImxlZ2FjeS10b2tlbi1rZXkiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJmY2I4NzZlM2NlNDI0MjZlODc0ODIwMDQxZjFmMWRkMy1yIiwic3ViIjoiMTA2NjFkNDctNmQwOC00ODcyLTgxYWEtNzg2ODMyN2NmZjBiIiwic2NvcGUiOlsib3BlbmlkIl0sImlhdCI6MTUzMDI3MTAzNiwiZXhwIjoxNTMyODYzMDM2LCJjaWQiOiJvbW5pYXV0aC1leGFtcGxlIiwiY2xpZW50X2lkIjoib21uaWF1dGgtZXhhbXBsZSIsImlzcyI6Imh0dHBzOi8vMTkyLjE2OC41MC42Ojg0NDMvb2F1dGgvdG9rZW4iLCJ6aWQiOiJ1YWEiLCJncmFudF90eXBlIjoiYXV0aG9yaXphdGlvbl9jb2RlIiwidXNlcl9uYW1lIjoiZHJuaWMiLCJvcmlnaW4iOiJ1YWEiLCJ1c2VyX2lkIjoiMTA2NjFkNDctNmQwOC00ODcyLTgxYWEtNzg2ODMyN2NmZjBiIiwicmV2X3NpZyI6Ijg0MmQ5N2RkIiwiYXVkIjpbIm9wZW5pZCIsIm9tbmlhdXRoLWV4YW1wbGUiXX0.ky1_EzIE0uVW2jNHeAS6crKIACM4LmN53IHUaHqT3acwUKc4aj_Ob2kClbPIZqeBycYe59I7zEfu5A0BwwcZy2ZrSaEGS34h8JmTqDLvgIDusz2DkaHHZ9xSeZooi9czLih-1JMMJVDZFbr_thtZiq6UqRdZTvpCMPTiF6pS8Pu0BYhjGFt2TZOuq9RkSMRC1L026HCmPAE-Z44UlWFj5mJ-B4nC9qeJKg4Doc05I8WV7_rWI_7yyDTyJ4W0A1k-18SyS_gb7aVfD0BWbu_PKQjmmx7XPHwcxlaSw0xcBcbJlQ7BVS4l3_GcjEZYYiZirAuhNAf8JHEBmxC_SVkpnA","authorized_scopes":"openid"},"extra":{"raw_info":{"user_id":"10661d47-6d08-4872-81aa-7868327cff0b","sub":"10661d47-6d08-4872-81aa-7868327cff0b","user_name":"drnic","given_name":"Dr Nic","family_name":"Williams","email":"drnic@starkandwayne.com","previous_logon_time":1530168890029,"name":"Dr Nic Williams"}}}
```

## Docker

```text
docker build -t starkandwayne/uaa-example-omniauth-login-only .
docker run -ti -p 9292:9292 -e UAA_URL=$UAA_URL -e UAA_CA_CERT=$UAA_CA_CERT starkandwayne/uaa-example-omniauth-login-only
```