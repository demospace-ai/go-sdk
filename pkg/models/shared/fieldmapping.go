package shared

type FieldMapping struct {
	DestinationFieldName *string `json:"destination_field_name,omitempty"`
	SourceFieldName      *string `json:"source_field_name,omitempty"`
}
