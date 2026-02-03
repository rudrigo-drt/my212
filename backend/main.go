package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	apiKeyUsername := os.Getenv("API_KEY_USERNAME")
	apiKeyPassword := os.Getenv("API_KEY_PASSWORD")

	if apiKeyUsername == "" {
		panic("TRADING212_API_KEY_USERNAME not defined")
	}

	if apiKeyPassword == "" {
		panic("TRADING212_API_KEY_PASSWORD not defined")
	}

	reqUrl := "https://live.trading212.com/api/v0/equity/account/summary"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(apiKeyUsername, apiKeyPassword)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println("\n" + string(body))
}
