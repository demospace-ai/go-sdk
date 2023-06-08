# Object

## Overview

Operations on objects

### Available Operations

* [CreateObject](#createobject) - Create a new object
* [GetObjects](#getobjects) - Get all objects

## CreateObject

Create a new object

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
    res, err := s.Object.CreateObject(ctx, shared.ObjectInput{
        CursorField: fabra.String("updated_at"),
        DestinationID: 2,
        DisplayName: "BigQuery",
        EndCustomerIDField: "end_customer_id",
        Frequency: 30,
        FrequencyUnits: shared.FrequencyUnitsHours,
        Namespace: "bigquery_dataset",
        ObjectFields: []shared.ObjectField{
            shared.ObjectField{
                Name: fabra.String("event_name"),
                Type: shared.FieldTypeJSON.ToPointer(),
            },
            shared.ObjectField{
                Name: fabra.String("event_name"),
                Type: shared.FieldTypeInteger.ToPointer(),
            },
            shared.ObjectField{
                Name: fabra.String("event_name"),
                Type: shared.FieldTypeTimestamp.ToPointer(),
            },
        },
        PrimaryKey: fabra.String("event_id"),
        TableName: "events",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateObject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetObjects

Get all objects

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Object.GetObjects(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetObjects200ApplicationJSONObject != nil {
        // handle response
    }
}
```
