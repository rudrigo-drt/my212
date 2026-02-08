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

func (t212 *Trading212Client) GetExchangesMetadata() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/metadata/exchanges"

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

func (t212 *Trading212Client) GetAvailableInstruments() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/metadata/instruments"

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

func (t212 *Trading212Client) GetPendingOrders(id string) (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/orders"
	if id != "" {
		reqUrl += "/" + id
	}

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

func (t212 *Trading212Client) GetOpenPositions(ticker string) (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/positions"

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return "", err
	}

	if ticker != "" {
		query := req.URL.Query()
		query.Add("ticker", ticker)
		req.URL.RawQuery = query.Encode()
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

func (t212 *Trading212Client) GetPies(id string) (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/pies"
	if id != "" {
		reqUrl += "/" + id
	}

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

func (t212 *Trading212Client) GetPaidOutDividends() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/history/dividends"

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return "", err
	}

	/*query := req.URL.Query()
	query.Add("cursor", "0")
	query.Add("ticker", "string")
	query.Add("limit", "21")
	req.URL.RawQuery = query.Encode()*/

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

func (t212 *Trading212Client) GetOrdersHistory() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/history/orders"

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

func (t212 *Trading212Client) GetTransactionsHistory() (string, error) {
	reqUrl := "https://live.trading212.com/api/v0/equity/history/transactions"

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return "", err
	}

	/*query := req.URL.Query()
	query.Add("cursor", "0")
	query.Add("time", "2026-01-01T00:00:00Z")
	query.Add("limit", "21")
	req.URL.RawQuery = query.Encode()*/

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
