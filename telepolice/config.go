package telepolice

type Config struct {
	Concurrency int
}

func NewConfig(concurrency int) *Config {
	return &Config{
		Concurrency: concurrency,
	}
}
