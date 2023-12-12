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
        BigqueryConfig: &shared.BigQueryConfig{
            Credentials: gosdk.String("Paste JSON from GCP"),
            Location: "us-west1",
        },
        ConnectionType: shared.ConnectionTypeSnowflake,
        DisplayName: "BigQuery",
        MongodbConfig: &shared.MongoDbConfig{
            ConnectionOptions: gosdk.String("retryWrites=true&w=majority"),
            Host: "examplecluster.abc123.mongodb.net",
            Password: "securePassword123",
            Username: "jane_doe",
        },
        RedshiftConfig: &shared.RedshiftConfig{
            DatabaseName: "your_database",
            Host: "examplecluster.12345.us-west-1.redshift.amazonaws.com",
            Password: "securePassword123",
            Port: "5432",
            Username: "jane_doe",
        },
        SnowflakeConfig: &shared.SnowflakeConfig{
            DatabaseName: "your_database",
            Host: "abc123.us-east4.gcp.snowflakecomputing.com",
            Password: "securePassword123",
            Role: "your_role",
            Username: "jane_doe",
            WarehouseName: "your_warehouse",
        },
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
| sdkerrors.SDKError | 400-600            | */*                |

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
| sdkerrors.SDKError | 400-600            | */*                |
