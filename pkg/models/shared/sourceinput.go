package shared

type SourceInput struct {
	BigqueryConfig  *BigQueryConfig    `json:"bigquery_config,omitempty"`
	ConnectionType  ConnectionTypeEnum `json:"connection_type"`
	DisplayName     string             `json:"display_name"`
	EndCustomerID   int64              `json:"end_customer_id"`
	MongodbConfig   *MongoDbConfig     `json:"mongodb_config,omitempty"`
	RedshiftConfig  *RedshiftConfig    `json:"redshift_config,omitempty"`
	SnowflakeConfig *SnowflakeConfig   `json:"snowflake_config,omitempty"`
}
