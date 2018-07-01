# CLI password client for Airports API filtered by UAA scopes

In another terminal, ensure that the Airports API app is running at https://localhost:9292.

In your current terminal, setup `$UAA_URL`/`$UAA_CA_CERT`:

```text
source <(path/to/uaa-deployment/bin/uaa-deployment env)
```


When you run the CLI you will be prompted for an Airports/UAA user/password:

```text
bundle install
bundle exec cli.rb
```

Enter the user credentials for `airports-no-scope`, with password `airports-no-scope`, and it will fetch the 20 results it is qualified to view from the API (as a user without an `airports.50` nor `airports.all` scope):

```text
username > airports-no-scope
password > *****************
+-------------------------------------+------+
| Name                                | ICAO |
+-------------------------------------+------+
| Brisbane Archerfield Airport        | YBAF |
| Northern Peninsula Airport          | YBAM |
| Alice Springs Airport               | YBAS |
| Brisbane International Airport      | YBBN |
| Gold Coast Airport                  | YBCG |
| Cairns International Airport        | YBCS |
| Charleville Airport                 | YBCV |
| Mount Isa Airport                   | YBMA |
| Sunshine Coast Airport              | YBMC |
| Mackay Airport                      | YBMK |
| Proserpine Whitsunday Coast Airport | YBPN |
| Rockhampton Airport                 | YBRK |
| Townsville Airport                  | YBTL |
| Weipa Airport                       | YBWP |
| Avalon Airport                      | YMAV |
| Albury Airport                      | YMAY |
| Melbourne Essendon Airport          | YMEN |
| RAAF Base East Sale                 | YMES |
| Hobart International Airport        | YMHB |
| Launceston Airport                  | YMLT |
+-------------------------------------+------+
```

The username and password are not stored. Instead, the CLI only stores the access token/refresh token/token type etc.

Our CLI stores them in `.user.json`:

```text
cat .user.json
```

To "log out", delete this file, and run the CLI to login as `airports-50` or `airports-all` user:

```text
$ bundle exec cli.rb
username > airports-all
password > ************
+-------------------------------------+------+
| Name                                | ICAO |
+-------------------------------------+------+
| Brisbane Archerfield Airport        | YBAF |
| Northern Peninsula Airport          | YBAM |
| Alice Springs Airport               | YBAS |
| Brisbane International Airport      | YBBN |
| Gold Coast Airport                  | YBCG |
...297 results...
```