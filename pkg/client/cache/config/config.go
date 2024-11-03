package config

import "time"

// RedisConfig is interface for redis config settings
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	ElementTTL() time.Duration
}
