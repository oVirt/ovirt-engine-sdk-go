# oVirt Engine API Go SDK

## Introduction

This project contains the Go SDK for the oVirt Engine API.

## Important

Note that most of the code of this SDK is automatically generated. You 
cloud just use this repository then you will have everything alread,
but if you want to use SDK locally, you can downloaded the source and 
generate it, please follow the instructions in the `README.md` file of 
the parent directory.

> Current version of Go SDK just use basic auth when accessing oVirt engine.

> Async api operation is currently not supported.

## Installation

The SDK can be installed in any operating systems with Go installed. Then
do the following:
```bash
$ go get -u github.com/imjoey/ovirt-engine-sdk-go/sdk/ovirtsdk4
```

## Usage

To use the SDK you should import ovirtsdk4 package as follows:
```go
import (
    "github.com/imjoey/ovirt-engine-sdk-go/sdk/ovirtsdk4"
)
```

That will give you access to all the classes of the SDK, and in particular
to the `Connection` class. This is the entry point of the SDK,
and gives you access to the root of the tree of services of the API:

```go
// Create a connection to the server:
import (
    "time"
    "github.com/imjoey/ovirt-engine-sdk-go/sdk/ovirtsdk4"
)

var inputRawURL = "https://engine.example.com/ovirt-engine/api"

conn, err := NewConnectionBuilder().
	URL(inputRawURL).
	Username("your-username").
	Password("your-password").
	Insecure(true).
	Compress(true).
	Timeout(time.Second * 10).
	Build()

// Get the reference to the **system** service
systemService := conn.SystemService()

// Get the reference to the **clusters** service
clustersService := systemService.ClustersService()

// Get all the clusters
clusters, err := clustersService.List(false, false, 100, "", nil, nil, false)

// Print clusters attrs
for _, cluster := range clusters {
    fmt.Printf("cluster(%v): CPU architecture is %v and type is %v", *cluster.Id,
        cluster.Cpu.Architecture, *cluster.Cpu.Type)
}

// Close the connection
conn.Close()
```

More usage examples will be added to the `examples` directory soon.
