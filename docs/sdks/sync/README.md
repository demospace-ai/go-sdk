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
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	gosdk "github.com/fabra-io/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(""),
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
        Namespace: gosdk.String("end_customer_bigquery_dataset"),
        ObjectID: 3,
        PrimaryKey: gosdk.String("event_id"),
        SourceID: 1,
        TableName: gosdk.String("end_customer_events"),
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
| `request`                                                | [shared.SyncInput](../../pkg/models/shared/syncinput.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateSyncResponse](../../pkg/models/operations/createsyncresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSyncs

Get all syncs

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
        gosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Sync.GetSyncs(ctx)
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

**[*operations.GetSyncsResponse](../../pkg/models/operations/getsyncsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
