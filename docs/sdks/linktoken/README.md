# LinkToken
(*LinkToken*)

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
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LinkToken.CreateLinkToken(ctx, shared.CreateLinkTokenRequest{
        EndCustomerID: "abcd-1234-efgh-5678",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateLinkTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateLinkTokenRequest](../../models/shared/createlinktokenrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateLinkTokenResponse](../../models/operations/createlinktokenresponse.md), error**

