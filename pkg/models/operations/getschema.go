package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetSchemaQueryParams struct {
	ConnectionID int64  `queryParam:"style=form,explode=true,name=connectionID"`
	Namespace    string `queryParam:"style=form,explode=true,name=namespace"`
	TableName    string `queryParam:"style=form,explode=true,name=table_name"`
}

type GetSchemaRequest struct {
	QueryParams GetSchemaQueryParams
}

type GetSchema200ApplicationJSON struct {
	Schema []shared.ColumnSchema `json:"schema,omitempty"`
}

type GetSchemaResponse struct {
	ContentType                       string
	StatusCode                        int
	RawResponse                       *http.Response
	GetSchema200ApplicationJSONObject *GetSchema200ApplicationJSON
}
