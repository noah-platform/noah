package main

import (
	"fmt"

	"github.com/noah-platform/noah/account/server/client"
)

const userID = "bank"

func main() {
	client := client.New(client.Config{
		BaseURL:  "http://localhost:8080",
		RetryMax: 3,
	})

	account, err := client.FetchAccount(userID)
	if err != nil {
		panic(err)
	}

	fmt.Print(account)
}
