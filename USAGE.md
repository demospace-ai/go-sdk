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
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetNamespacesRequest{
        ConnectionID: 548814,
    }

    res, err := s.Connection.GetNamespaces(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Namespaces != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->