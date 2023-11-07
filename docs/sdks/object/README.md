# Object
(*.Object*)

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
        gosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Object.CreateObject(ctx, shared.ObjectInput{
        CursorField: gosdk.String("updated_at"),
        DestinationID: 2,
        DisplayName: "BigQuery",
        EndCustomerIDField: "end_customer_id",
        Frequency: 30,
        FrequencyUnits: shared.FrequencyUnitsMinutes,
        Namespace: "bigquery_dataset",
        ObjectFields: []shared.ObjectField{
            shared.ObjectField{
                Name: gosdk.String("event_name"),
            },
        },
        PrimaryKey: gosdk.String("event_id"),
        TableName: "events",
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
        gosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Object.GetObjects(ctx)
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

**[*operations.GetObjectsResponse](../../models/operations/getobjectsresponse.md), error**

