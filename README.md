# Example applications using Cloud Foundry UAA

This repository attempts to implement a series of similar example applications, in different programming languages.

Currently, the applications are implmented in:

* [Ruby](/ruby)
* [Go](/golang)

The sample applications:

* Resource Server - A simple API that provides a list of Australian airports. Guests only receive 10 results. Authenticated users receive more. Users who have specific UAA claims/scopes can receive all the results.
* Resource Server CLI - A CLI for interacting with the above Resource Server.
* Resource Server Web UI - A frontend wrapper UI for the backend Resource Server: a map showing Australian airports. Guest users see only 10 airports. Once a user has been authenticated via your UAA, they will begin to see more airports. Add the user to the `airports.all` group and they will be authorized to see all the airports.

## Dependencies on pull requests

These various applications or tutorials require the following PRs to be merged and/or releases cut:

* https://github.com/cloudfoundry-incubator/uaa-cli/pull/12 - `uaa context --auth_header`
* https://github.com/cloudfoundry-incubator/uaa-cli/issues/4 - `uaa` version cut + binaries released
* https://github.com/cloudfoundry/cf-uaa-lib/pull/31 - README shows full example including decode using `/token_keys`
* https://github.com/cloudfoundry/cf-uaa-lib/pull/32 - upgrade aesthetics of hashes to use ruby2 {key: value} syntax; add ruby 2.5.1 to travis
* https://github.com/cloudfoundry/cf-uaa-lib/pull/33 - explicit `WARNING` for skipping verification of token signing; or if token claims it wasn't signed
* https://github.com/cloudfoundry/cf-uaa-lib/pull/34 - explicit `TokenCoder.decode_token_expiry`
* https://github.com/cloudfoundry/cf-uaa-lib/pull/35 - Automatically look up /token_keys to decode access token
* https://github.com/cloudfoundry/cf-uaa-lib/pull/37 - `TokenCoder.decode` + `TokenCoder#decode` accept either access token or auth header as input
* https://github.com/cloudfoundry/cf-uaa-lib/pull/36 - UAA no longer has a /varz endpoint
* https://github.com/cloudfoundry/uaa/pull/868 - UAA no longer has a /varz endpoint