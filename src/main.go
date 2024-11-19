package main

import (
	"context"
	"fmt"
	"openfga-poc/src/open-fga"
)

func main() {
	config := GetConfig()

	client, err := open_fga.NewOpenFGAClient(config.OpenFgaUrl)
	if err != nil {
		panic(fmt.Sprintf("couldn't create OpenFGA client: %v", err))
	}

	siteId := "site"
	userId := "vitor"
	action := "report_publish"

	check, err := client.Check(context.Background(), siteId, userId, action)
	if err != nil {
		return
	}

	fmt.Printf("Is %s allowed to %s in %s? %t\n", userId, action, siteId, *check)
}
