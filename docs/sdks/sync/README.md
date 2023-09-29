# Sync
(*Sync*)

## Overview

Operations on syncs

### Available Operations

* [CreateSync](#createsync) - Create a new sync
* [GetSyncs](#getsyncs) - Get all syncs

## CreateSync

Create a new sync

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
    res, err := s.Sync.CreateSync(ctx, shared.SyncInput{
        CursorField: gosdk.String("updated_at"),
        CustomJoin: gosdk.String("select * from events join additional_properties on events.id = additional_properties.event_id;"),
        DestinationID: 2,
        DisplayName: "Event Sync",
        EndCustomerID: "abc123",
        FieldMappings: []shared.FieldMapping{
            shared.FieldMapping{
                DestinationFieldName: gosdk.String("event"),
                SourceFieldName: gosdk.String("event_name"),
            },
        },
        Frequency: gosdk.Int64(30),
        FrequencyUnits: shared.FrequencyUnitsMinutes.ToPointer(),
        Namespace: gosdk.String("end_customer_bigquery_dataset"),
        ObjectID: 3,
        PrimaryKey: gosdk.String("event_id"),
        SourceID: 1,
        TableName: gosdk.String("end_customer_events"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSync200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.SyncInput](../../models/shared/syncinput.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateSyncResponse](../../models/operations/createsyncresponse.md), error**


## GetSyncs

Get all syncs

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
    res, err := s.Sync.GetSyncs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSyncs200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSyncsResponse](../../models/operations/getsyncsresponse.md), error**

