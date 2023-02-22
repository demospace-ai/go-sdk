package shared

type SnowflakeConfig struct {
	DatabaseName  string `json:"database_name"`
	Host          string `json:"host"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Username      string `json:"username"`
	WarehouseName string `json:"warehouse_name"`
}
