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
+-------------------------------------+------+----------+
| Name                                | ICAO | Altitude |
+-------------------------------------+------+----------+
| Brisbane Archerfield Airport        | YBAF | 63       |
| Northern Peninsula Airport          | YBAM | 34       |
| Alice Springs Airport               | YBAS | 1789     |
| Brisbane International Airport      | YBBN | 13       |
| Gold Coast Airport                  | YBCG | 21       |
| Cairns International Airport        | YBCS | 10       |
| Charleville Airport                 | YBCV | 1003     |
| Mount Isa Airport                   | YBMA | 1121     |
| Sunshine Coast Airport              | YBMC | 15       |
| Mackay Airport                      | YBMK | 19       |
| Proserpine Whitsunday Coast Airport | YBPN | 82       |
| Rockhampton Airport                 | YBRK | 34       |
| Townsville Airport                  | YBTL | 18       |
| Weipa Airport                       | YBWP | 63       |
| Avalon Airport                      | YMAV | 35       |
| Albury Airport                      | YMAY | 539      |
| Melbourne Essendon Airport          | YMEN | 282      |
| RAAF Base East Sale                 | YMES | 23       |
| Hobart International Airport        | YMHB | 13       |
| Launceston Airport                  | YMLT | 562      |
+-------------------------------------+------+----------+
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
+-------------------------------------+------+----------+
| Name                                | ICAO | Altitude |
+-------------------------------------+------+----------+
| Brisbane Archerfield Airport        | YBAF | 63       |
| Northern Peninsula Airport          | YBAM | 34       |
| Alice Springs Airport               | YBAS | 1789     |
| Brisbane International Airport      | YBBN | 13       |
| Gold Coast Airport                  | YBCG | 21       |
...297 results...
```