package config

type Config struct {
}

type DBConfig struct {
	Host          string   `json:"host"`
	ReadOnlyHosts []string `json:"read_only_hosts"`
	Port          uint     `json:"port"`
	User          string   `json:"user"`
	Password      string   `json:"password"`
	DBName        string   `json:"db_name"`
	TZ            string   `json:"tz"`
}
