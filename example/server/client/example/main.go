package main

import (
	"fmt"

	"github.com/noah-platform/noah/example/server/client"
)

const exampleID = "1"

func main() {
	client := client.New(client.Config{
		BaseURL:  "http://localhost:8080",
		RetryMax: 3,
	})

	example, err := client.FetchExample(exampleID)
	if err != nil {
		panic(err)
	}

	fmt.Print(example)
}
