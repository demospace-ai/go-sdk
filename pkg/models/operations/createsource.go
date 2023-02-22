package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type CreateSourceRequest struct {
	Request shared.SourceInput `request:"mediaType=application/json"`
}

type CreateSource200ApplicationJSON struct {
	Source *shared.Source `json:"source,omitempty"`
}

type CreateSourceResponse struct {
	ContentType                          string
	StatusCode                           int
	CreateSource200ApplicationJSONObject *CreateSource200ApplicationJSON
}
