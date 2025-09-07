package pkg

import (
	"net/http"
	"time"

	"log"

	"golang.org/x/net/proxy"
)

func NewTorClient() *http.Client {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Ошибка создания Tor SOCKS5: %v", err)
	}

	transport := &http.Transport{
		Dial:               dialer.Dial,
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}
