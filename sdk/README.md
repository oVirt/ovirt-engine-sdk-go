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

inputRawURL := "https://10.1.111.229/ovirt-engine/api"
conn, err := NewConnectionBuilder().
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

// Get the reference to the "clusters" service:
clustersService := conn.SystemService().ClustersService()

// Use the "list" method of the "clusters" service to list all the clusters of the system:
clustersResponse, err := clustersService.List().Send()
if err != nil {
	fmt.Printf("Failed to get cluster list, reason: %v\n", err)
	return
}

if clusters, ok := clustersResponse.Clusters(); ok {
	// Print the datacenter names and identifiers:
	fmt.Printf("Cluster: (")
	for _, cluster := range clusters {
		if clusterName, ok := cluster.Name(); ok {
			fmt.Printf(" name: %v", clusterName)
		}
		if clusterId, ok := cluster.Id(); ok {
			fmt.Printf(" id: %v", clusterId)
		}
	}
	fmt.Println(")")
}
```

You could find more usage examples in `examples` directory.
