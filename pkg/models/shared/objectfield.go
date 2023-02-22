package shared

type ObjectField struct {
	Name *string        `json:"name,omitempty"`
	Type *FieldTypeEnum `json:"type,omitempty"`
}
