package shared

type Connection struct {
	ConnectionType *ConnectionTypeEnum `json:"connection_type,omitempty"`
	ID             *int64              `json:"id,omitempty"`
}
