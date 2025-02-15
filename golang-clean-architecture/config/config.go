package config

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Cert     string
	Driver   string
}

type APIConfig struct {
	Port string
}

type Config struct {
	DBConfig
	APIConfig
}

func (c *Config) loadConfig() error {

	c.DBConfig = DBConfig{
		Host:     "unthenamed-unthenamed.h.aivencloud.com",
		Port:     13689,
		User:     "avnadmin",
		Password: "AVNS_DilIFOqbMOlCyx0DADG",
		Database: "el_db",
		Cert:     "ca.pm",
		Driver:   "postgres",
	}

	c.APIConfig = APIConfig{
		Port: "8080",
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.loadConfig()

	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}
