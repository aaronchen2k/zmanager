package model

type Config struct {
	ZTFVersion string
	ZDVersion  string

	Interval int64
	Language string
}

func NewConfig() Config {
	return Config{
		Interval: 30,
		Language: "en",
	}
}
