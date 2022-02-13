package serverapi

type Config struct {
	Addr string `yaml:"host"`
}

func NewConfig() *Config {
	return &Config{
		Addr: ":8080",
	}
}
