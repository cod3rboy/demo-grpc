# gRPC Demo (Invoice generation and query)

A demo gRPC application consisting of a server and client which deals in the generation and query of transaction invoices.

## Start Server

`make run-server`

or

`go run ./cmd/server`

> **Note:** Pass optional `-port <number>` flag to the above `go run` command to specify a different port. Default port is `8000`.

## Run Client

### Create Invoice

`make invoice-create`

or

`go run ./cmd/client --action=create`

> **Note:** Pass optional `-server <host:port>` flag to above `go run` command to specify different server address. Default address is `localhost:8000`.

### Query Invoice

`make invoice-query`

or

`go run ./cmd/client --action=query`

> **Note:** Pass optional `-server <host:port>` flag to above `go run` command to specify different server address. Default address is `localhost:8000`.
