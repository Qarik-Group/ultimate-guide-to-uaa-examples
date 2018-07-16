# Example applications using Cloud Foundry UAA

These various applications or tutorials require the following PRs to be merged and/or releases cut:

* https://github.com/cloudfoundry/omniauth-uaa-oauth2/pull/12 - fix deprecation warning
* https://github.com/cloudfoundry/omniauth-uaa-oauth2/pull/14 - support custom root CA
* https://github.com/cloudfoundry/omniauth-uaa-oauth2/issues/13 - cut new release
* https://github.com/cloudfoundry-incubator/uaa-cli/pull/12 - `uaa context --auth_header`
* https://github.com/cloudfoundry-incubator/uaa-cli/issues/4 - `uaa` version cut + binaries released
* https://github.com/cloudfoundry/cf-uaa-lib/pull/31 - README shows full example including decode using `/token_keys`
* https://github.com/cloudfoundry/cf-uaa-lib/pull/32 - upgrade aesthetics of hashes to use ruby2 {key: value} syntax; add ruby 2.5.1 to travis
* https://github.com/cloudfoundry/cf-uaa-lib/pull/33 - explicit `WARNING` for skipping verification of token signing; or if token claims it wasn't signed
* https://github.com/cloudfoundry/cf-uaa-lib/pull/34 - explicit `TokenCoder.decode_token_expiry`
* https://github.com/cloudfoundry/cf-uaa-lib/pull/35 - Automatically look up /token_keys to decode access token

* https://github.com/cloudfoundry/cf-uaa-lib/pull/36 - UAA no longer has a /varz endpoint
* https://github.com/cloudfoundry/uaa/pull/868 - UAA no longer has a /varz endpoint