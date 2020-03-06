package configs

import (
	"time"
)

type ConfigDatabase struct {
	URL     string
	Timeout time.Duration
}

func NewConfig(url string, timeout time.Duration) ConfigDatabase {
	return ConfigDatabase{
		URL:     url,
		Timeout: timeout,
	}
}
