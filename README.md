# oVirt Go SDK
Go SDK for oVirt 4.0+, main functions are finished, excluding error processing...

Project completion is almost 80%.

## Introduction

The oVirt Go SDK is a Go package that simplyfies access to the
oVirt Engine API.

## Building

You must install the Go binary and setup the Go environments, including
`GOROOT` and `GOPATH`.

Most of the source code of the Go SDK is automatically generated
from the API model (Java).

The code generator is a Java program that resides in the `generator`
directory. This Java program will get the API model and the metamodel
artifacts (offered by oVirt team) from the available Maven repositories. 
To build and run it use the following commands:

```bash
$ git clone git@github.com:imjoey/ovirt-engine-sdk-go.git
$ cd ovirt-engine-sdk-go
$ mvn package
```

This will build the code generator, run it to generate the SDK for the
version of the API that corresponds to the branch of the SDK that you
are using.

If you need to generate it for a different version of the API then you
can use the `model.version` property. For example, if you need to
generate the SDK for version `4.1.0` of the SDK you can use this
command:
```bash
$ mvn package -Dmodel.version=4.1.0
```

By default the build and the tests are executed using the `go` command.
If you wish to use a different version of Go you can use the
`go.command` property:
```bash
$ mvn package -Dgo.command=go1.7
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

For more usage examples, you could refer to [sdk/README](https://github.com/imjoey/ovirt-engine-sdk-go/blob/master/sdk/README.md).
