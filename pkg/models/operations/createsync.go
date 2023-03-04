package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSyncRequest struct {
	Request shared.SyncInput `request:"mediaType=application/json"`
}

type CreateSync200ApplicationJSON struct {
	Sync *shared.Sync `json:"sync,omitempty"`
}

type CreateSyncResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	CreateSync200ApplicationJSONObject *CreateSync200ApplicationJSON
}
