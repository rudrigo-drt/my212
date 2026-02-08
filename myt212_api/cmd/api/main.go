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
		fmt.Print("\n> ")
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
		case "2":
			res, err := t212.GetExchangesMetadata()
			if err != nil {
				fmt.Println("Error getting Exchanges Metadata: ", err)
				continue
			}
			fmt.Println(res)
		case "3":
			res, err := t212.GetAvailableInstruments()
			if err != nil {
				fmt.Println("Error getting Available Instruments: ", err)
				continue
			}
			fmt.Println(res)
		case "4":
			fmt.Print("Enter order ID (or leave empty for all): ")
			orderId, _ := reader.ReadString('\n')
			orderId = strings.TrimSpace(orderId)
			res, err := t212.GetPendingOrders(orderId)
			if err != nil {
				fmt.Println("Error getting Pending Order(s): ", err)
				continue
			}
			fmt.Println(res)
		case "5":
			fmt.Print("Enter ticker (or leave empty for all): ")
			ticker, _ := reader.ReadString('\n')
			ticker = strings.TrimSpace(ticker)
			res, err := t212.GetOpenPositions(ticker)
			if err != nil {
				fmt.Println("Error getting Open Position(s): ", err)
				continue
			}
			fmt.Println(res)
		case "6":
			fmt.Print("Enter pie ID (or leave empty for all): ")
			pieId, _ := reader.ReadString('\n')
			pieId = strings.TrimSpace(pieId)
			res, err := t212.GetPies(pieId)
			if err != nil {
				fmt.Println("Error getting Pie(s): ", err)
				continue
			}
			fmt.Println(res)
		case "7":
			res, err := t212.GetPaidOutDividends()
			if err != nil {
				fmt.Println("Error getting Paid Out Dividends: ", err)
				continue
			}
			fmt.Println(res)
		case "8":
			res, err := t212.GetOrdersHistory()
			if err != nil {
				fmt.Println("Error getting Orders History: ", err)
				continue
			}
			fmt.Println(res)
		case "9":
			res, err := t212.GetTransactionsHistory()
			if err != nil {
				fmt.Println("Error getting Transactions History: ", err)
				continue
			}
			fmt.Println(res)
		case "0", "quit", "exit":
			fmt.Println("Quiting...")
			return
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}
