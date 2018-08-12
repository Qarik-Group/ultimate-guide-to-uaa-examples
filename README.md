# Example applications using Cloud Foundry UAA

This repository attempts to implement a series of similar example applications, in different programming languages.

Currently, the applications are implmented in:

* [Ruby](/ruby)
* [Go](/golang)

The sample applications:

* Resource Server - A simple API that provides a list of Australian airports. Guests only receive 10 results. Authenticated users receive more. Users who have specific UAA claims/scopes can receive all the results.
* Resource Server CLI - A CLI for interacting with the above Resource Server.
* Resource Server Web UI - A frontend wrapper UI for the backend Resource Server: a map showing Australian airports. Guest users see only 10 airports. Once a user has been authenticated via your UAA, they will begin to see more airports. Add the user to the `airports.all` group and they will be authorized to see all the airports.

## Deploy UAA

You can run these example applications and integrate with any UAA.

The tutorials makes some small assumptions that you're using either [`uaa-deployment-cf`](https://github.com/starkandwayne/uaa-deployment-cf/) or [`uaa-deployment`](https://github.com/starkandwayne/uaa-deployment/) for deployment and setting up local env vars.

The former makes it easy to deploy a secure, production-grade UAA to your Cloud Foundry using `cf push`. The latter makes it easy to deploy a secure, production-grade UAA to any target cloud using `bosh create-env`. Both tools provide a `u` CLI with a matching set of subcommands. To deploy your UAA, install and setup the `uaa` CLI, and setup local environment variables for the example applications:

```text
cd ~/workspace/uaa-deployment*/
source <(bin/u env)
u up
source <(bin/u env)
u auth-client
```

The example applications' tutorials assume that either of the two `u up` projects are cloned into the `~/workspace/uaa-deployment-cf` or `~/workspace/uaa-deployment` folders respectively.

### BYO UAA

If you are not using one of these two `u up` tools for deploying your UAA, and do not have the `u env` helper to setup environment variables, then you will need to manually setup the following env vars:

* `$UAA_URL` - the URL for your UAA, such as https://login.mycompany.com
* `$UA_CA_CERT` - [optional] the custom root CA for your UAA URL
* `$UA_CA_CERT_FILE` - [optional] the custom root CA for your UAA URL stored in a local file

You also need to install the `uaa` CLI.

Wherever the instructions say `u auth-client` you will need to manually run the following commands to target and authenticate:

```text
uaa target https://login.mycompany.com
uaa get-client-credentials-token uaa_admin --client_secret uaa_admin_secret
```

In the example above, the `uaa_admin` client should be replaced with an existing UAA client that has the `uaa.admin` authority to allow you to create new clients, users, groups, etc.

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