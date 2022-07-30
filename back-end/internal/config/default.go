package config

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/database/mongo"
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/handler"
	"github.com/amirhnajafiz/personal-website/back-end/internal/jwt"
)

const (
	tokenTimeout = 10
)

func Default() Config {
	return Config{
		Address: 8080,
		Http: handler.Config{
			Admin: handler.Admin{
				User: "",
				Pass: "",
			},
		},
		JWT: jwt.Config{
			Key:     "",
			Timeout: tokenTimeout,
		},
		Mongodb: mongo.Config{
			Database: "",
			URL:      "",
		},
	}
}
