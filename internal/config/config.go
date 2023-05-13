package config

type Config struct {
	Name        string
	Host        string            `env:"app_host"`
	Port        int               `env:"app_port"`
	Environment map[string]string `env:"environment"`
	Volumes     []string          `env:"volumes"`
}

func NewConfig() *Config {
	return &Config{
		Name:        "Street",
		Host:        "localhost",
		Port:        8080,
		Environment: map[string]string{"key": "value"},
		Volumes:     []string{"volume1", "volume2"},
	}
}
