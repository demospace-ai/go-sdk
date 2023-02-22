package shared

type Object struct {
	CustomerIDColumn *string       `json:"customer_id_column,omitempty"`
	DestinationID    *int64        `json:"destination_id,omitempty"`
	DisplayName      *string       `json:"display_name,omitempty"`
	ID               *int64        `json:"id,omitempty"`
	Namespace        *string       `json:"namespace,omitempty"`
	ObjectFields     []ObjectField `json:"object_fields,omitempty"`
	TableName        *string       `json:"table_name,omitempty"`
}
