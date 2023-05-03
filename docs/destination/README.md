# Destination

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
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Destination.CreateDestination(ctx, shared.DestinationInput{
        BigqueryConfig: &shared.BigQueryConfig{
            Credentials: fabra.String("Paste JSON from GCP"),
            Location: "us-west1",
        },
        ConnectionType: shared.ConnectionTypeEnumWebhook,
        DisplayName: "BigQuery",
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

    if res.CreateDestination200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetDestinations

Get all destinations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Destination.GetDestinations(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDestinations200ApplicationJSONObject != nil {
        // handle response
    }
}
```
