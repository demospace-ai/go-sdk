package shared

type ColumnSchema struct {
	Name *string        `json:"name,omitempty"`
	Type *FieldTypeEnum `json:"type,omitempty"`
}
