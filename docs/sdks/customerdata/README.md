# CustomerData
(*CustomerData*)

### Available Operations

* [QueryObject](#queryobject) - Query object record for customer

## QueryObject

Query a single object record directly from a customer's source. The response payload will match the object schema you've defined.

### Example Usage

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
    var endCustomerID string = "once"
    var objectID int64 = 819181
    requestBody := &operations.QueryObjectRequestBody{
        Filters: []shared.QueryFilter{
            shared.QueryFilter{
                FieldName: "user_id",
                FieldValue: "2",
            },
        },
    }

    ctx := context.Background()
    res, err := s.CustomerData.QueryObject(ctx, endCustomerID, objectID, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.QueryObject200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `endCustomerID`                                                                         | *string*                                                                                | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `objectID`                                                                              | *int64*                                                                                 | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `requestBody`                                                                           | [*operations.QueryObjectRequestBody](../../models/operations/queryobjectrequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |


### Response

**[*operations.QueryObjectResponse](../../models/operations/queryobjectresponse.md), error**

