package main

import (
	"fmt"
	"io"
	"net/http"

	"main/internal/config"
)

func main() {
	cfg := config.Load()

	reqUrl := "https://live.trading212.com/api/v0/equity/account/summary"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(cfg.APIKeyUsername, cfg.APIKeyPassword)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println("\n" + string(body))
}
