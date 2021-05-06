package twdataloader

import (
	"net/http"
	"time"
)

func getDefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}
