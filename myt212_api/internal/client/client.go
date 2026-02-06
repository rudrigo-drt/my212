package client

import (
	"io"
	"log"
	"net/http"
	"time"

	"myt212_api/internal/config"
)

type Trading212Client struct {
	Cfg        *config.Config
	httpClient *http.Client
}

func New(cfg *config.Config) *Trading212Client {
	return &Trading212Client{
		Cfg: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (t212 *Trading212Client) GetAccountSummary() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/account/summary"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return "", err
	}

	log.Println(req.Method, req.URL)

	req.SetBasicAuth(t212.Cfg.APIKeyUsername, t212.Cfg.APIKeyPassword)
	res, err := t212.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	log.Println("Status: ", res.Status)

	return string(body), nil
}
