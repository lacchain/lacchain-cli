## LACChain CLI

This is a command line client to operate a LACChain node in the LACChain Besu Network.

## Prerequisites

* Go 1.12+ installation or later
* **GOPATH** environment variable is set correctly
* docker version 17.03 or later

## Package overview

1. **cmd** contains the commands
2. **lib** contains most of the code
4. **business** contains business logic that will be consume by APIs
5. **models** conatins data models of requests and responses of APIs
7. **util** contains util functions about files and ethereum address

## Install

```
$ git clone https://github.com/lacchain/lacchain-cli

$ export GO111MODULE=on

$ cd LacchainCli
$ go build -o $GOPATH/bin/lacchain
```

## Run

```
$ lacchain version
$ Lacchain Version 0.0.9
```
