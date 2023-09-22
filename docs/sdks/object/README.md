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
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Object.CreateObject(ctx, shared.ObjectInput{
        CursorField: gosdk.String("updated_at"),
        DestinationID: 2,
        DisplayName: "BigQuery",
        EndCustomerIDField: "end_customer_id",
        Frequency: 30,
        FrequencyUnits: shared.FrequencyUnitsHours,
        Namespace: "bigquery_dataset",
        ObjectFields: []shared.ObjectField{
            shared.ObjectField{
                Name: gosdk.String("event_name"),
                Type: shared.FieldTypeTimestamp.ToPointer(),
            },
        },
        PrimaryKey: gosdk.String("event_id"),
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.ObjectInput](../../models/shared/objectinput.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateObjectResponse](../../models/operations/createobjectresponse.md), error**


## GetObjects

Get all objects

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(shared.Security{
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetObjectsResponse](../../models/operations/getobjectsresponse.md), error**

