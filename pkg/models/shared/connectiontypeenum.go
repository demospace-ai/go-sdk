package shared

type ConnectionTypeEnum string

const (
	ConnectionTypeEnumSnowflake ConnectionTypeEnum = "snowflake"
	ConnectionTypeEnumBigquery  ConnectionTypeEnum = "bigquery"
	ConnectionTypeEnumRedshift  ConnectionTypeEnum = "redshift"
	ConnectionTypeEnumMongodb   ConnectionTypeEnum = "mongodb"
	ConnectionTypeEnumWebhook   ConnectionTypeEnum = "webhook"
)
