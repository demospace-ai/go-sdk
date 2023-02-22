package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type GetDestinations200ApplicationJSON struct {
	Destinations []shared.Destination `json:"destinations,omitempty"`
}

type GetDestinationsResponse struct {
	ContentType                             string
	StatusCode                              int
	GetDestinations200ApplicationJSONObject *GetDestinations200ApplicationJSON
}
