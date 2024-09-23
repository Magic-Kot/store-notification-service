package config

import "time"

type Config struct {
	GRPC       GRPCConfig `env:"GRPC" yaml:"grpc"`
	LoggerDeps LoggerDeps `env:"LOGGER" yaml:"logger"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type LoggerDeps struct {
	LogLevel string `env:"LOG_LEVEL" yaml:"logLevel" env-default:"info"`
}
