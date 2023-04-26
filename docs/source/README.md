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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.SourceInput{
        BigqueryConfig: &shared.BigQueryConfig{
            Credentials: fabra.String("Paste JSON from GCP"),
            Location: "us-west1",
        },
        ConnectionType: shared.ConnectionTypeEnumBigquery,
        DisplayName: "Frontend Events",
        EndCustomerID: 123,
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
    }

    res, err := s.Source.CreateSource(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetSources

Get all sources

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
    res, err := s.Source.GetSources(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSources200ApplicationJSONObject != nil {
        // handle response
    }
}
```
