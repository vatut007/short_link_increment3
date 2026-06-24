package config

type Config struct {
	Port    int
	Address string
}

var ConfigApp = Config{
	Port:    flagRunPort,
	Address: flagRunAddr,
}
