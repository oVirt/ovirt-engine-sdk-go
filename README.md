# oVirt Go SDK 
[![Build, Test, Release](https://github.com/oVirt/ovirt-engine-sdk-go/actions/workflows/build.yaml/badge.svg)](https://github.com/oVirt/ovirt-engine-sdk-go/actions/workflows/build.yaml)
[![CodeQL](https://github.com/oVirt/ovirt-engine-sdk-go/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/oVirt/ovirt-engine-sdk-go/actions/workflows/github-code-scanning/codeql)
[![Go Report Card](https://goreportcard.com/badge/github.com/oVirt/ovirt-engine-sdk-go)](https://goreportcard.com/report/github.com/oVirt/ovirt-engine-sdk-go)


## Introduction

The oVirt Go SDK is a Go package that simplifies access to the
oVirt Engine API.

> IMPORTANT: This document describes how to generate, build and test the
SDK. If you are interested in how to use it read the `README.md` file
in the [oVirt/go-ovirt](https://github.com/oVirt/go-ovirt) repository instead.

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
$ git clone git@github.com:oVirt/ovirt-engine-sdk-go.git
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
