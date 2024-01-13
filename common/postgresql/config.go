package postgresql

type Config struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime string
}
