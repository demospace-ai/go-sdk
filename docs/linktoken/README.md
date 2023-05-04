# LinkToken

## Overview

Operations on link tokens

### Available Operations

* [CreateLinkToken](#createlinktoken) - Create a new link token

## CreateLinkToken

Create a new link token

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
    res, err := s.LinkToken.CreateLinkToken(ctx, shared.CreateLinkTokenRequest{
        EndCustomerID: "123",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateLinkTokenResponse != nil {
        // handle response
    }
}
```
