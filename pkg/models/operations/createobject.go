package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
)

type CreateObjectRequest struct {
	Request shared.ObjectInput `request:"mediaType=application/json"`
}

type CreateObject200ApplicationJSON struct {
	Object *shared.Object `json:"object,omitempty"`
}

type CreateObjectResponse struct {
	ContentType                          string
	StatusCode                           int
	CreateObject200ApplicationJSONObject *CreateObject200ApplicationJSON
}
