package config

import "time"

type Config struct {
	ServerDeps  ServerDeps  `env:"SERVER" yaml:"server"`
	LoggerDeps  LoggerDeps  `env:"LOGGER" yaml:"logger"`
	MailService MailService `yaml:"mail_service"`
}

type ServerDeps struct {
	Host    string        `env:"HOST"  yaml:"host" env-default:"localhost"`
	Port    string        `env:"PORT" yaml:"port" env-default:"44044"`
	Timeout time.Duration `env:"TIMEOUT" yaml:"timeout" env-default:"5s"`
}

type LoggerDeps struct {
	LogLevel string `env:"LOG_LEVEL" yaml:"logLevel" env-default:"info"`
}

type MailService struct {
	Username string `env:"USERNAME_NS" yaml:"username"`
	Password string `env:"PASSWORD_NS" yaml:"password"`
}
