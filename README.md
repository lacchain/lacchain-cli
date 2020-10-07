## LACChain CLI

This is a command line client to operate a LACChain node in the LACChain Besu Network.

## Prerequisites

* Go 1.12+ installation or later
* **GOPATH** environment variable is set correctly

## Package overview

1. **cmd** contains the commands
2. **model** conatins data models of requests and responses of APIs

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
## Copyright 2020 LACChain

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.