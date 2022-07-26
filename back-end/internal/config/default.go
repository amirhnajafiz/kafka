package config

import "github.com/amirhnajafiz/personal-website/back-end/internal/database/mongo"

func Default() Config {
	return Config{
		Address: 8080,
		Mongodb: mongo.Config{
			URL: "",
		},
	}
}
