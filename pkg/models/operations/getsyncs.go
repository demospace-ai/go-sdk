package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type GetSyncs200ApplicationJSON struct {
	Syncs []shared.Sync `json:"syncs,omitempty"`
}

type GetSyncsResponse struct {
	ContentType                      string
	StatusCode                       int
	GetSyncs200ApplicationJSONObject *GetSyncs200ApplicationJSON
}
