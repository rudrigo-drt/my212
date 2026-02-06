package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"myt212_api/internal/client"
	"myt212_api/internal/config"
)

func main() {
	cfg := config.Load()
	t212 := client.New(&cfg)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			res, err := t212.GetAccountSummary()
			if err != nil {
				fmt.Println("Error getting Account Summary: ", err)
				continue
			}
			fmt.Println(res)
		case "quit", "exit":
			fmt.Println("Quiting")
			return
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}
