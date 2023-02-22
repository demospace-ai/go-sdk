package shared

type Destination struct {
	Connection  *Connection `json:"connection,omitempty"`
	DisplayName *string     `json:"display_name,omitempty"`
	ID          *int64      `json:"id,omitempty"`
}
