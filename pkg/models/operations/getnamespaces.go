package operations

type GetNamespacesQueryParams struct {
	ConnectionID int64 `queryParam:"style=form,explode=true,name=connectionID"`
}

type GetNamespacesRequest struct {
	QueryParams GetNamespacesQueryParams
}

type GetNamespaces200ApplicationJSON struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

type GetNamespacesResponse struct {
	ContentType                           string
	StatusCode                            int
	GetNamespaces200ApplicationJSONObject *GetNamespaces200ApplicationJSON
}
