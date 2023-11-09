# Connection
(*Connection*)

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
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(""),
    )


    var connectionID int64 = 995455

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

**[*operations.GetNamespacesResponse](../../pkg/models/operations/getnamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSchema

Get schema for table

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


    var connectionID int64 = 367941

    var namespace string = "string"

    var tableName string = "string"

    ctx := context.Background()
    res, err := s.Connection.GetSchema(ctx, connectionID, namespace, tableName)
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
| `connectionID`                                        | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `namespace`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `tableName`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetSchemaResponse](../../pkg/models/operations/getschemaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTables

Get all tables

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


    var connectionID int64 = 820803

    var namespace string = "string"

    ctx := context.Background()
    res, err := s.Connection.GetTables(ctx, connectionID, namespace)
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
| `connectionID`                                        | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `namespace`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetTablesResponse](../../pkg/models/operations/gettablesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
