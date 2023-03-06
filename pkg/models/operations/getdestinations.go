package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetDestinations200ApplicationJSON struct {
	Destinations []shared.Destination `json:"destinations,omitempty"`
}

type GetDestinationsResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	GetDestinations200ApplicationJSONObject *GetDestinations200ApplicationJSON
}
