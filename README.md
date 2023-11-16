# Fabra Go SDK

<div align="left">
   <p>Use the Fabra API to build customer-facing data warehouse integrations to let your customers start sending data to your application. Unblock your sales pipeline in days, not months.</p>
   <a href="https://github.com/fabra-io/go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/fabra-io/go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://www.fabra.io/#Email-Hero"><img src="https://img.shields.io/static/v1?label=Docs&message=Sign Up&color=2ca47c&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/fabra-io/go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Connection](docs/sdks/connection/README.md)

* [GetNamespaces](docs/sdks/connection/README.md#getnamespaces) - Get all namespaces
* [GetSchema](docs/sdks/connection/README.md#getschema) - Get schema for table
* [GetTables](docs/sdks/connection/README.md#gettables) - Get all tables

### [CustomerData](docs/sdks/customerdata/README.md)

* [QueryObject](docs/sdks/customerdata/README.md#queryobject) - Query object record for customer

### [Destination](docs/sdks/destination/README.md)

* [CreateDestination](docs/sdks/destination/README.md#createdestination) - Create a new destination
* [GetDestinations](docs/sdks/destination/README.md#getdestinations) - Get all destinations

### [LinkToken](docs/sdks/linktoken/README.md)

* [CreateLinkToken](docs/sdks/linktoken/README.md#createlinktoken) - Create a new link token

### [Object](docs/sdks/object/README.md)

* [CreateObject](docs/sdks/object/README.md#createobject) - Create a new object
* [GetObjects](docs/sdks/object/README.md#getobjects) - Get all objects

### [Source](docs/sdks/source/README.md)

* [CreateSource](docs/sdks/source/README.md#createsource) - Create a new source
* [GetSources](docs/sdks/source/README.md#getsources) - Get all sources

### [Sync](docs/sdks/sync/README.md)

* [CreateSync](docs/sdks/sync/README.md#createsync) - Create a new sync
* [GetSyncs](docs/sdks/sync/README.md#getsyncs) - Get all syncs
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/sdkerrors"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(""),
	)

	var connectionID int64 = 995455

	ctx := context.Background()
	res, err := s.Connection.GetNamespaces(ctx, connectionID)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.fabra.io` | None |

#### Example

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
		gosdk.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		gosdk.WithServerURL("https://api.fabra.io"),
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
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKeyAuth` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gosdk "github.com/fabra-io/go-sdk"
	"log"
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
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
