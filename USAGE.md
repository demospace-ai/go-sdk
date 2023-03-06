<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/fabra-io/go-sdk"
    "github.com/fabra-io/go-sdk/pkg/models/shared"
    "github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(fabra.WithSecurity(
        shared.Security{
            APIKeyAuth: shared.SchemeAPIKeyAuth{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
    ))
    
    req := operations.GetNamespacesRequest{
        QueryParams: operations.GetNamespacesQueryParams{
            ConnectionID: 548814,
        },
    }

    ctx := context.Background()
    res, err := s.Connection.GetNamespaces(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNamespaces200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->