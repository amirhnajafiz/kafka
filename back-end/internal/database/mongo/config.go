package mongo

type Config struct {
	Database string `konaf:"database"`
	URL      string `koanf:"url"`
}
