# go-ps-rest

> process list web server

`ps-rest` uses [procfs](https://en.wikipedia.org/wiki/Procfs) to obtain a list
of processes for display over a http api

## Installation

Running `go get github.com/razic/go-ps-rest` will install a `go-ps-rest` binary
to `$GOPATH/bin`.

## Supported Platforms

* Linux

## Usage

Run the binary. It doesn't take any flags. It will start an insecure HTTP
server bound to all interfaces on `localhost:8080`. The endpoint for the
process listing is `localhost:8080/ps`

### Start the server

```bash
$ go-ps-rest
```

### List all processes

```bash
curl localhost:8080/ps
```
