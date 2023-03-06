package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type CreateLinkTokenRequest struct {
	Request shared.CreateLinkTokenRequest `request:"mediaType=application/json"`
}

type CreateLinkTokenResponse struct {
	ContentType             string
	CreateLinkTokenResponse *shared.CreateLinkTokenResponse
	StatusCode              int
	RawResponse             *http.Response
}
