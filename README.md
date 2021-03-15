# Elastic Cloud Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/elastic/cloud-sdk-go.svg)](https://pkg.go.dev/github.com/elastic/cloud-sdk-go)

Go SDK for Elastic Cloud. Its goal is to provide common ground for all Elastic Cloud programmatic code in Go.

## Installation

Run the following `go get` command to install the SDK in your module dependencies directory:

```console
go get -u github.com/elastic/cloud-sdk-go
```

## Usage

See the [`pkg/api`](https://pkg.go.dev/github.com/elastic/cloud-sdk-go/pkg/api) package for more in depth documentation.

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
)

var (
	logFormat = log.Lmsgprefix | log.Llongfile

	errLog  = log.New(os.Stderr, "ERROR ", logFormat)
	warnLog = log.New(os.Stdout, "WARN ", logFormat)
	infoLog = log.New(os.Stdout, "INFO ", logFormat)
)

func main() {
	// Export your apikey as an environment variable as EC_API_KEY. To generate
	// a new API key go to ESS or ECE Web UI > API Keys > Generate API Key.
	apiKey := os.Getenv("EC_API_KEY")
	if apiKey == "" {
		warnLog.Print("unable to obtain value from EC_API_KEY environment variable")
	}

	// Create a API instance with an API key as means of authentication.
	ess, err := api.NewAPI(api.Config{
		Client:     new(http.Client),
		AuthWriter: auth.APIKey(apiKey),
	})
	if err != nil {
		errLog.Fatal(err)
	}

	// List the user's deployments via the `deploymentapi` package (Recommended).
	res, err := deploymentapi.List(deploymentapi.ListParams{API: ess})
	if err != nil {
		errLog.Fatal(err)
	}
	infoLog.Printf("found %d deployents", len(res.Deployments))

	encoder := json.NewEncoder(infoLog.Writer())
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(res); err != nil {
		errLog.Fatal(err)
	}
}
```

## High level package overview

The project's structure is based off the standard Go project layout. Therefore, all of our library code that we expect other projects to import is placed in the `pkg/` directory.

The main packages for interacting directly with our public API endpoints can be found within the [`pkg/api`](https://github.com/elastic/cloud-sdk-go/tree/master/pkg/api) directory. The source code for these APIs are the [`client`](https://github.com/elastic/cloud-sdk-go/tree/master/pkg/client) and [`models`](https://github.com/elastic/cloud-sdk-go/tree/master/pkg/models), which are generated off the [public API swagger specification](./api/apidocs.json).

The rest are a series of packages that can be leveraged in many ways. For a detailed package description visit the SDK's packages documentation at [pkg.go.dev](https://pkg.go.dev/github.com/elastic/cloud-sdk-go). Alternatively you can use the [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) command on the root level of this project.

## Getting started

[Ecctl](https://github.com/elastic/ecctl) (the Elastic Cloud CLI tool) depends heavily on cloud-sdk-go. You can use ecctl command packages as a reference on how to leverage the SDK. One good example is the [command](https://github.com/elastic/ecctl/blob/master/cmd/deployment/list.go) used to list deployments.

Alternatively, if you wish to write your own APIs, take a look at our [sample code](./examples/platform/main.go), which retrieves information about the active platform, to get an idea on how the client is used.

We always welcome contributions! Take a look at our [contributing guide](./CONTRIBUTING.md) if this is something that interests you.
