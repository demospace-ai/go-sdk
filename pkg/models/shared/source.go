package shared

type Source struct {
	Connection    *Connection `json:"connection,omitempty"`
	DisplayName   *string     `json:"display_name,omitempty"`
	EndCustomerID *int64      `json:"end_customer_id,omitempty"`
	ID            *int64      `json:"id,omitempty"`
}
