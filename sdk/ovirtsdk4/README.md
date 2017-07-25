# oVirt Go SDK [![Build Status](https://travis-ci.org/imjoey/ovirt-engine-sdk-go.svg?branch=master)](https://travis-ci.org/imjoey/ovirt-engine-sdk-go)

## Introduction

The oVirt Go SDK is a Go package that simplyfies access to the
oVirt Engine API.

> IMPORTANT: The code in this project is generated automatically by the [imjoey/ovirt-engine-sdk-go](https://github.com/imjoey/ovirt-engine-sdk-go). So if you want to know how to generate the code, please read the `README.md` in the  [imjoey/ovirt-engine-sdk-go](https://github.com/imjoey/ovirt-engine-sdk-go) repositorys instead.

## Usage

To use the SDK you should import ovirtsdk4 package as follows:

```go
import (
    "github.com/imjoey/go-ovirt"
)
```

That will give you access to all the classes of the SDK, and in particular
to the `Connection` class. This is the entry point of the SDK,
and gives you access to the root of the tree of services of the API:

```go
// Create a connection to the server:
import (
    "time"
    "github.com/imjoey/go-ovirt"
)

inputRawURL := "https://10.1.111.229/ovirt-engine/api"
conn, err := ovirtsdk4.NewConnectionBuilder().
	URL(inputRawURL).
	Username("admin@internal").
	Password("qwer1234").
	Insecure(true).
	Compress(true).
	Timeout(time.Second * 10).
	Build()
if err != nil {
	t.Fatalf("Make connection failed, reason: %s", err.Error())
}

defer conn.Close()

clustersListResponse, err2 := conn.SystemService().ClustersService().
	List().
	CaseSensitive(false).
	Max(100).
	Send()

if err2 != nil {
	t.Fatalf("Get clusters failed, reason: %s", err2.Error())
}

for _, cluster := range clustersListResponse.Clusters() {
	t.Logf("cluster(%v): CPU architecture is %v and type is %v", *cluster.Id,cluster.Cpu.Architecture, *cluster.Cpu.Type)
}

```
