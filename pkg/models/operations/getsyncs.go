package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetSyncs200ApplicationJSON struct {
	Syncs []shared.Sync `json:"syncs,omitempty"`
}

type GetSyncsResponse struct {
	ContentType                      string
	StatusCode                       int
	RawResponse                      *http.Response
	GetSyncs200ApplicationJSONObject *GetSyncs200ApplicationJSON
}
