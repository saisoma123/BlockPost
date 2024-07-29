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
4. Run `go mod tidy` to install the Go packages this app uses
5. `make npm_build` command will install all of the necessary JS packages to run the
React app.
6. `sudo apt install jq` for transaction querying.
7. Also install konsole with `sudo apt install konsole`.
8. `make` command will start the app, it runs `ignite chain serve` to start the blockchain itself,
then it starts a JS server and React app that you can interact with.
9. `make runAllTests` command will run all of the tests for the main components of the app,
the msg_server, query_server, keeper state management. To run individual tests,
use these commands, `make keeper_test`, `make query_server_test`, `make msg_server_test`

Note that this is a localhost app, so the design decisions I made are secure for
this particular app, as it is entirely local to the user. But, if I was to deploy this to an actual network, I would add
wallet integration, and have the querier be an actual blockchain transaction that
the user has to sign off on. This is so that my app would be able to verify that the stored 
message creator address is the same address as the wallet calling for the 
transaction. This would make it so that only the user who stored the original message
can retrieve it and no one else.
