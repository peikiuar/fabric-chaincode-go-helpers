# fabric-chaincode-go-helpers

![Lint Code Base](https://github.com/braduf/fabric-chaincode-go-helpers/workflows/Lint%20Code%20Base/badge.svg)
![Test Code Base](https://github.com/braduf/fabric-chaincode-go-helpers/workflows/Test%20Code%20Base/badge.svg)

## Introduction

Libraries with general helper functions that are often needed and used in the development of Hyperledger Fabric Chaincode in Go. By using these packages you can focus on your smart contract's business logic and forget about underlying Chaincode boilerplate code.

## How to use

Since this is a private repository, we need two extra steps to use these packages in other projects. 

The first step is telling Go that the repo it needs to get is private:

```shell
go env -w GOPRIVATE=github.com/peikiuar/fabric-chaincode-go-helpers
```

And since Go with Go modules uses git to get the imported packages in a project, the second step is to make sure that our local git configurations have our credentials to access this repository. This can be set with the following command:

```shell
git config \
  --global \
  url."https://${user}:${personal_access_token}@github.com".insteadOf \
  "https://github.com"
```

After that the this module or packages can be imported normally in any project that uses Go modules. For example, to use the state package:

```go
import "github.com/peikiuar/fabric-chaincode-go-helpers/state"
```
