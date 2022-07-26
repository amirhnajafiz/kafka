package config

import "github.com/amirhnajafiz/personal-website/back-end/internal/database/mongo"

func Default() Config {
	return Config{
		Mongo: mongo.Config{
			URL: "",
		},
	}
}
