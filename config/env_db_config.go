package config

import "os"

type EnvDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewEnvDBConfig() (*EnvDBConfig, error) {
	return &EnvDBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}, nil
}

func (c *EnvDBConfig) GetHost() string {
	return c.Host
}

func (c *EnvDBConfig) GetPort() string {
	return c.Port
}

func (c *EnvDBConfig) GetUser() string {
	return c.User
}

func (c *EnvDBConfig) GetPassword() string {
	return c.Password
}

func (c *EnvDBConfig) GetDBName() string {
	return c.DBName
}
