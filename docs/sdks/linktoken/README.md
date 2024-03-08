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
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	gosdk "github.com/fabra-io/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateLinkTokenRequest](../../pkg/models/shared/createlinktokenrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateLinkTokenResponse](../../pkg/models/operations/createlinktokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
