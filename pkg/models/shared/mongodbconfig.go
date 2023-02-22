package shared

type MongoDbConfig struct {
	ConnectionOptions *string `json:"connection_options,omitempty"`
	Host              string  `json:"host"`
	Password          string  `json:"password"`
	Username          string  `json:"username"`
}
