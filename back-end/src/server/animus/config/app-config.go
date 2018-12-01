package config

import (
	"os"
	"server/app/config"
)

type AppConfig struct {

}

func NewAppConfig() config.IAppConfig {
	return &AppConfig{}
}

func (a *AppConfig) TokenPassword() string {
	return os.Getenv("token_password")
}

func (a *AppConfig) Port() string {
	return os.Getenv("port")
}