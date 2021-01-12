package model

type Config struct {
	ZTFVersion float64
	ZDVersion  float64

	Interval int64
	Language string
}

func NewConfig() Config {
	return Config{
		Interval: 6,
		Language: "en",
	}
}
