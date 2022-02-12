package serverapi

type Config struct {
	BindAddr string `yaml:"host"` /*todo: change name*/
	/*Store *store.Config*/
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		/*Store: store.NewConfig(),*/
	}
}
