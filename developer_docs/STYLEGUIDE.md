# Style Guide

- [General patterns](#general-patterns)
- [Errors](#errors)
  - [Error handling](#error-handling)
  - [API Errors](#api-errors)
  - [Multiple errors](#multiple-errors)
- [Packages](#packages)
  - [Structure](#structure)
  - [Util packages](#util-packages)
- [Testing](#testing)
  - [Unit tests](#unit-tests)
- [General Style](#general-style)
  - [Usage of pointers](#usage-of-pointers)
  - [Arrays and slices](#arrays-and-slices)
  - [Config or params as the only function parameter](#config-or-params-as-the-only-function-parameter)
  - [Functions and methods](#functions-and-methods)
- [Documentation](#documentation)
- [Extras](#extras)

## General patterns

Our codebase is structured very similarly to what [DDD (Domain Driven Design)](https://martinfowler.com/bliki/EvansClassification.html) dictates, and even though there’s no strict naming or patterns being enforced we try to follow it as a good practice.

## Errors

All errors should be informative and helpful. Something like "deployment ID not specified and is required for this operation" is more useful than "no ID". This is a helpful blog post on the subject: [How to Write Good Error Messages.](https://uxplanet.org/how-to-write-good-error-messages-858e4551cd4)

Error strings should not be capitalised or end with punctuation, since they are usually printed following other context.

In regards to packages, we prefer to use the `"errors"` standard library package over `"github.com/pkg/errors"`

### Error handling

All errors must be handled and returned, `_` variables must not be used to discard them.

### API Errors

API errors should always be encapsulated with [`apierror.Unwrap()`](https://github.com/elastic/cloud-sdk-go/blob/master/pkg/api/apierror/unwrap.go#L59), this function tries to break down and inspect the encapsulated and multi-layer wraps that the API errors contain.

### Multiple errors

When multiple errors can be returned, it is preferable to use the [`mutlierror.Prefixed`](https://github.com/elastic/cloud-sdk-go/blob/master/pkg/multierror/prefixed.go#L41) type to return all the possible errors with a prefixed string to include some context.

yes! :smile:

```go
func (params StopParams) Validate() error {
    var merr = multierror.NewPrefixed("invalid deployment stop params")
    if params.ID == "" {
        merr = merr.Append(errors.New("at least 1 instance ID must be provided"))
    }

    if params.API == nil {
        merr = merr.Append(apierror.ErrMissingAPI)
    }

   return merr.ErrorOrNil()
}
```

preferably not :confused:

```go
func (params StopParams) Validate() error {
    if params.ID == "" {
        return errors.New("ID cannot be empty")
    }

    if params.API == nil {
        return errors.New("api reference is required")
    }

   return nil
}
```

## Packages

### Structure

A package per context, a context applies to any of the high level containers like `platform`, `deployment`, etc. When a context becomes too large to be contained in a package we can start breaking it down into sub-packages.

API packages are defined following the structure of the Elastic Cloud API. If a request to the API is `GET /deployments/templates/{template_id}`, the corresponding package will be [/pkg/api/deploymentapi/deptemplateapi](../pkg/api/deploymentapi/deptemplateapi).

### Util packages

When a function can be made generic it should go in one of the utils packages (e.g. [pkg/util](../pkg/util)) to remove complexity and give the ability to be reused.

## Testing

### Unit tests

All files containing functions or methods must have a corresponding unit test file, and we aim to have 100% coverage.

#### API Mocks

When unit testing functions which will call the external API, please use the provided `api.NewMock` in conjunction with a mock response assertion function.

yes! :smile:

``` go
//Test case
{
    name: "succeeds",
    args: args{params{
        API: api.NewMock(mock.New200ResponseAssertion(
            &mock.RequestAssertion{
                Header: api.DefaultWriteMockHeaders,
                Method: "DELETE",
                Host:   api.DefaultMockHost,
                Path:   "/api/v1/deployments/templates/some-id",
                Query: url.Values{
                    "region": []string{"us-east-1"},
                },
            },
            mock.NewStringBody(`{}`),
        )),
    }},
},
// More tests ...
```

## General Style

Before committing to your feature branch, run `make lint` and `make format` to ensure the code has the proper styling format. We run linters on the CI build, so this also prevents your builds from failing.

### Usage of pointers

Unless a pointer is necessary, always use the value type before switching to the pointer type. Of course, if your structure contains a mutex, or needs to be synced between different goroutines, it needs to be a pointer, otherwise there’s no reason why it should be.

### Arrays and slices

To remove complexity, when it's not necessary to set length or capacity, it's preferable to declare a nil slice as it has no underlying array. We should avoid using the `make` function, as this function allocates a zeroed array that returns a slice of said array.

yes! :smile:

```go
var slice []string
```

preferably not :confused:

```go
slice := make([]string, 0)
```

### `Config` or `Params` as the only function parameter

A `Config` or `Params` structure is used to encapsulate all the parameters in a structure that has a `.Validate()` error signature, so it can be validated inside the receiver of that structure.

Unless the `Params` struct needs to satisfy the `pool.Validator` interface for concurrency on that receiver it should always remain a value type and not a pointer type.

### Functions and methods

Names should be descriptive and we should avoid redundancy.

yes! :smile:

```go
deploymentapi.Create()
```

preferably not :confused:

```go
deploymentapi.CreateDeployment()
```

When using method chaining make sure to put each method in it's own line to improve readability.

yes! :smile:

```go
res, res2, err := params.V1API.Deployments.CreateDeployment(
    deployments.NewCreateDeploymentParams().
        WithRequestID(id).
        WithBody(params.Request),
    params.AuthWriter,
)
```

preferably not :confused:

```go
res, res2, err := params.V1API.Deployments.CreateDeployment(
    deployments.NewCreateDeploymentParams().WithRequestID(id).WithBody(params.Request),
    params.AuthWriter,
)
```

When possible we try to avoid `else` and nested `if`s. This makes our code more readable and removes complexity.
Specifically, we prefer having multiple `return` statements over having nested code.

yes! :smile:

``` go
if params.Hide {
    params.Do()
    return data, nil
}

if isHidden {
    return nil, fmt.Errorf("example error", params.Name)
}
return data, nil
```

preferably not :confused:

``` go
if params.Hide {
        params.Do()
    } else {
        if isHidden {
            return nil, fmt.Errorf("example error", params.Name)
        }
    }
```

## Documentation

The package wide description and documentation is provided in a godoc `doc.go` file. Aside form packages with a very small context, all packages should have this file.

## Extras

For further information on good practices with Go in general, check out this [document](https://github.com/golang/go/wiki/CodeReviewComments).
