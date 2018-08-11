# Example client applications in Go

This folder contains a series of example client applications. Each subfolder contains one application and a progressive tutorial in its README. The subsequent folders assume that you have completed the preceding application tutorials.

The sample applications:

* [Resource Server](resource-server/) - A simple API that provides a list of Australian airports. Guests only receive 10 results. Authenticated users receive more. Users who have specific UAA claims/scopes can receive all the results.
* [Resource Server CLI](resource-server-cli/) - A CLI for interacting with the above Resource Server.
* [Resource Server Web UI](resource-server-wrapper-ui/) - A frontend wrapper UI for the backend Resource Server: a map showing Australian airports. Guest users see only 10 airports. Once a user has been authenticated via your UAA, they will begin to see more airports. Add the user to the `airports.all` group and they will be authorized to see all the airports.
