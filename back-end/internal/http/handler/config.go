package handler

type Config struct {
	Admin Admin `koanf:"admin"`
}

type Admin struct {
	User string `konaf:"user"`
	Pass string `konaf:"pass"`
}
