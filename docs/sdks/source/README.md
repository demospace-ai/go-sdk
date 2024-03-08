# Source
(*Source*)

## Overview

Operations on sources

### Available Operations

* [CreateSource](#createsource) - Create a new source
* [GetSources](#getsources) - Get all sources

## CreateSource

Create a new source

### Example Usage

```go
package main

import(
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	gosdk "github.com/fabra-io/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Source.CreateSource(ctx, shared.SourceInput{
        ConnectionType: shared.ConnectionTypeWebhook,
        DisplayName: "Frontend Events",
        EndCustomerID: "abcd-1234-efgh-5678",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.SourceInput](../../pkg/models/shared/sourceinput.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.CreateSourceResponse](../../pkg/models/operations/createsourceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSources

Get all sources

### Example Usage

```go
package main

import(
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	gosdk "github.com/fabra-io/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Source.GetSources(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSourcesResponse](../../pkg/models/operations/getsourcesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
