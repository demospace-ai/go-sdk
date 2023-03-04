package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetObjects200ApplicationJSON struct {
	Objects []shared.Object `json:"objects,omitempty"`
}

type GetObjectsResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetObjects200ApplicationJSONObject *GetObjects200ApplicationJSON
}
