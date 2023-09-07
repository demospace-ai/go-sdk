# Source

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
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Source.CreateSource(ctx, shared.SourceInput{
        BigqueryConfig: &shared.BigQueryConfig{
            Credentials: fabra.String("Paste JSON from GCP"),
            Location: "us-west1",
        },
        ConnectionType: shared.ConnectionTypeRedshift,
        DisplayName: "Frontend Events",
        EndCustomerID: "abcd-1234-efgh-5678",
        MongodbConfig: &shared.MongoDbConfig{
            ConnectionOptions: fabra.String("retryWrites=true&w=majority"),
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

    if res.CreateSource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.SourceInput](../../models/shared/sourceinput.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateSourceResponse](../../models/operations/createsourceresponse.md), error**


## GetSources

Get all sources

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Source.GetSources(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSources200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSourcesResponse](../../models/operations/getsourcesresponse.md), error**

