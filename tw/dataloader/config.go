package dataloader

import (
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
	Client  *http.Client
}

func (cfg *Config) Init() {
	if cfg.Client == nil {
		cfg.Client = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
}
