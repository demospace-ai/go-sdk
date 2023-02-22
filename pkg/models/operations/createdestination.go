package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
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
	CreateDestination200ApplicationJSONObject *CreateDestination200ApplicationJSON
}
