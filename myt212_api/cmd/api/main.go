package main

import (
	"myt212_api/cmd/cli"
	"myt212_api/internal/client"
	"myt212_api/internal/config"
	"myt212_api/internal/server"
)

func main() {
	//1. Load configuration
	cfg := config.Load()

	//2. Start API server
	srv := server.New(cfg.Address)
	go srv.Start()

	//3. Initialize Trading212 API client
	t212 := client.New(&cfg)

	//4. Start Command Line Interface
	cli := cli.New(t212)
	cli.Run()
}
