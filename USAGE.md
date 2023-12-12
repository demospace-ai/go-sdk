<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
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
<!-- End SDK Example Usage [usage] -->