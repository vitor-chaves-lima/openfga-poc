package main

import (
	"context"
	"fmt"
)

func main() {
	config := GetConfig()

	client, err := NewOpenFGAClient(config.OpenFgaUrl)
	if err != nil {
		panic(fmt.Sprintf("couldn't create OpenFGA client: %v", err))
	}

	err = Initialize(context.Background(), client)
	if err != nil {
		panic(fmt.Sprintf("couldn't initialize OpenFGA store: %v", err))
	}
}
