# BlockPost
**BlockPost** is a blockchain app that allows for the storage and retrieval of messages.

## Get started

1. [Download and install Go](https://golang.org/dl/)
2. Add Go binary to your PATH and set up Go workspace:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```
3. Install the Ignite CLI: `go install github.com/ignite/cli/ignite@latest`
4. `make npm_build` command will install all of the necessary JS packages to run the
React app.
5. `make` command will start the app, it runs `ignite chain serve` to start the blockchain itself,
then it starts a JS server and React app that you can interact with.
6. `make runAllTests` command will run all of the tests for the main components of the app,
the msg_server, query_server, keeper state management. To run individual tests,
use these commands, `make keeper_test`, `make query_server_test`, `make msg_server_test` 