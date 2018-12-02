package config

type IAppConfig interface {
	TokenPassword() string
	Port() string
	DatabaseUrl() string
}
