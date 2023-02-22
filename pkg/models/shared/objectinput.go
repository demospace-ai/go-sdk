package shared

type ObjectInput struct {
	CustomerIDColumn string        `json:"customer_id_column"`
	DestinationID    int64         `json:"destination_id"`
	DisplayName      string        `json:"display_name"`
	Namespace        string        `json:"namespace"`
	ObjectFields     []ObjectField `json:"object_fields,omitempty"`
	TableName        string        `json:"table_name"`
}
