package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type GetSources200ApplicationJSON struct {
	Sources []shared.Source `json:"sources,omitempty"`
}

type GetSourcesResponse struct {
	ContentType                        string
	StatusCode                         int
	GetSources200ApplicationJSONObject *GetSources200ApplicationJSON
}
