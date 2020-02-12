package telepolice

type Config struct {
	Concurrency int
	// Pod immediately after startup is in preparation and passes health check for the specified number of seconds
	IgnorerablePodStartTimeOfSec int
}

func NewConfig(concurrency int, ignorerablePodStartTimeOfSec int) *Config {
	return &Config{
		Concurrency:                  concurrency,
		IgnorerablePodStartTimeOfSec: ignorerablePodStartTimeOfSec,
	}
}
