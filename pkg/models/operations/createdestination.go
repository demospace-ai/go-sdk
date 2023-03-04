package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationRequest struct {
	Request shared.DestinationInput `request:"mediaType=application/json"`
}

type CreateDestination200ApplicationJSON struct {
	Destination *shared.Destination `json:"destination,omitempty"`
}

type CreateDestinationResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	CreateDestination200ApplicationJSONObject *CreateDestination200ApplicationJSON
}
