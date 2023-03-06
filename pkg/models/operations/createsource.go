package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
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
	RawResponse                          *http.Response
	CreateSource200ApplicationJSONObject *CreateSource200ApplicationJSON
}
