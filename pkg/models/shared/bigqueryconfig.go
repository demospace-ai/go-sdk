package shared

type BigQueryConfig struct {
	Credentials *string `json:"credentials,omitempty"`
	Location    string  `json:"location"`
}
