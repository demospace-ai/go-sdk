package shared

type FieldTypeEnum string

const (
	FieldTypeEnumString    FieldTypeEnum = "string"
	FieldTypeEnumInteger   FieldTypeEnum = "integer"
	FieldTypeEnumTimestamp FieldTypeEnum = "timestamp"
	FieldTypeEnumJSON      FieldTypeEnum = "json"
	FieldTypeEnumBoolean   FieldTypeEnum = "boolean"
)
