# Destination
(*Destination*)

## Overview

Operations on destinations

### Available Operations

* [CreateDestination](#createdestination) - Create a new destination
* [GetDestinations](#getdestinations) - Get all destinations

## CreateDestination

Create a new destination

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
    res, err := s.Destination.CreateDestination(ctx, shared.DestinationInput{
        ConnectionType: shared.ConnectionTypeSnowflake,
        DisplayName: "BigQuery",
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.DestinationInput](../../pkg/models/shared/destinationinput.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateDestinationResponse](../../pkg/models/operations/createdestinationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDestinations

Get all destinations

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
    res, err := s.Destination.GetDestinations(ctx)
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

**[*operations.GetDestinationsResponse](../../pkg/models/operations/getdestinationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
