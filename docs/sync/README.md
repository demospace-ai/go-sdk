# Sync

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
    res, err := s.Sync.CreateSync(ctx, shared.SyncInput{
        CursorField: fabra.String("updated_at"),
        CustomJoin: fabra.String("select * from events join additional_properties on events.id = additional_properties.event_id;"),
        DestinationID: 2,
        DisplayName: "Event Sync",
        EndCustomerID: "abc123",
        FieldMappings: []shared.FieldMapping{
            shared.FieldMapping{
                DestinationFieldName: fabra.String("event"),
                SourceFieldName: fabra.String("event_name"),
            },
            shared.FieldMapping{
                DestinationFieldName: fabra.String("event"),
                SourceFieldName: fabra.String("event_name"),
            },
            shared.FieldMapping{
                DestinationFieldName: fabra.String("event"),
                SourceFieldName: fabra.String("event_name"),
            },
            shared.FieldMapping{
                DestinationFieldName: fabra.String("event"),
                SourceFieldName: fabra.String("event_name"),
            },
        },
        Frequency: fabra.Int64(30),
        FrequencyUnits: shared.FrequencyUnitsMinutes.ToPointer(),
        Namespace: fabra.String("end_customer_bigquery_dataset"),
        ObjectID: 3,
        PrimaryKey: fabra.String("event_id"),
        SourceID: 1,
        TableName: fabra.String("end_customer_events"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSync200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetSyncs

Get all syncs

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
    res, err := s.Sync.GetSyncs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSyncs200ApplicationJSONObject != nil {
        // handle response
    }
}
```
