package jwt

import "time"

type Config struct {
	Key     string        `koanf:"key"`
	Timeout time.Duration `koanf:"timeout"`
}
