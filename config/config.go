package config

type Configuration struct {
	Env      string
	Database Database
	HTTP     HTTP
}

type Database struct {
	Name            string
	User            string
	Password        string
	Host            string
	Port            string
	ConnectionRetry int
}

type HTTP struct {
	Domain string
	Port   string
}
