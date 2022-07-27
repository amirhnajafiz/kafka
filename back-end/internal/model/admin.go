package model

type Admin struct {
	User string `koanf:"user"`
	Pass string `koanf:"pass"`
}
