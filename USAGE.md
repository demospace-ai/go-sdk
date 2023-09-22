<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    connectionID := 548814

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
<!-- End SDK Example Usage -->