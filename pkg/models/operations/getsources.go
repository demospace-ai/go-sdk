package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetSources200ApplicationJSON struct {
	Sources []shared.Source `json:"sources,omitempty"`
}

type GetSourcesResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetSources200ApplicationJSONObject *GetSources200ApplicationJSON
}
