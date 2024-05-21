package config

type Configuration struct {
	GoEnv    *GoEnv
	Database *Database
	HTTP     *HTTP
}

type GoEnv struct {
	Env string
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
