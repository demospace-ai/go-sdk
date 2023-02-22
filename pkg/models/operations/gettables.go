package operations

type GetTablesQueryParams struct {
	ConnectionID int64  `queryParam:"style=form,explode=true,name=connectionID"`
	Namespace    string `queryParam:"style=form,explode=true,name=namespace"`
}

type GetTablesRequest struct {
	QueryParams GetTablesQueryParams
}

type GetTables200ApplicationJSON struct {
	Tables []string `json:"tables,omitempty"`
}

type GetTablesResponse struct {
	ContentType                       string
	StatusCode                        int
	GetTables200ApplicationJSONObject *GetTables200ApplicationJSON
}
