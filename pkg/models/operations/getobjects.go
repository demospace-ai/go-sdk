package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type GetObjects200ApplicationJSON struct {
	Objects []shared.Object `json:"objects,omitempty"`
}

type GetObjectsResponse struct {
	ContentType                        string
	StatusCode                         int
	GetObjects200ApplicationJSONObject *GetObjects200ApplicationJSON
}
