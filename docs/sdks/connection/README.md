# Connection

## Overview

Operations on connections

### Available Operations

* [GetNamespaces](#getnamespaces) - Get all namespaces
* [GetSchema](#getschema) - Get schema for table
* [GetTables](#gettables) - Get all tables

## GetNamespaces

Get all namespaces

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    connectionID := 592845

    ctx := context.Background()
    res, err := s.Connection.GetNamespaces(ctx, connectionID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Namespaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `connectionID`                                        | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetNamespacesResponse](../../models/operations/getnamespacesresponse.md), error**


## GetSchema

Get schema for table

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    connectionID := 715190
    namespace := "quibusdam"
    tableName := "unde"

    ctx := context.Background()
    res, err := s.Connection.GetSchema(ctx, connectionID, namespace, tableName)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSchema200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `connectionID`                                        | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `namespace`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `tableName`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetSchemaResponse](../../models/operations/getschemaresponse.md), error**


## GetTables

Get all tables

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    connectionID := 857946
    namespace := "corrupti"

    ctx := context.Background()
    res, err := s.Connection.GetTables(ctx, connectionID, namespace)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTables200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `connectionID`                                        | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `namespace`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetTablesResponse](../../models/operations/gettablesresponse.md), error**
