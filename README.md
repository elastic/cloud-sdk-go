# Elastic Cloud Go SDK

Go SDK for Elastic Cloud. Its goal is to provide common ground for all Elastic Cloud programmatic code in Go.

## Installation

Run the following `go get` command to install the SDK in your module dependencies directory

```console
go get -u github.com/elastic/cloud-sdk-go
```

## High level package overview

The project's structure is based off the standard Go project layout. Therefore, all of our library code that we expect other projects to import is placed in the `pkg/` directory.

The main packages for interacting directly with our public API endpoints are `client` and `models`. The source code for these is generated off the [public API swagger specification](./api/apidocs.json]).

The rest are a series of packages that can be leveraged in many ways. For a detailed package description visit the SDK's packages documentation at [godoc.org](https://godoc.org/elastic/cloud-sdk-go). Alternatively you can use the [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command on the root level of this project.

## Getting started

Take a look at our [sample code](./examples/platform/main.go), which retrieves information about the active platform, to get an idea on how to use the SDK.

ecctl (the Elastic Cloud CLI tool) depends heavily on cloud-sdk-go. You can use ecctl packages as a reference on how to leverage the SDK. One good example is the [code](https://github.com/elastic/ecctl/blob/master/pkg/deployment/elasticsearch/list.go) used to list elasticsearch clusters.

We always welcome contributions! Take a look at our [contributing guide](./CONTRIBUTING.md) if this is something that interests you.
