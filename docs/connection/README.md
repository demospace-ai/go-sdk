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
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Connection.GetNamespaces(ctx, 592845)
    if err != nil {
        log.Fatal(err)
    }

    if res.Namespaces != nil {
        // handle response
    }
}
```

## GetSchema

Get schema for table

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Connection.GetSchema(ctx, 715190, "quibusdam", "unde")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSchema200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetTables

Get all tables

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Connection.GetTables(ctx, 857946, "corrupti")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTables200ApplicationJSONObject != nil {
        // handle response
    }
}
```
