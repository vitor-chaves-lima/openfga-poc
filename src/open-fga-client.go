package main

import (
	"context"
	"fmt"
	"log"

	. "github.com/openfga/go-sdk/client"
)

func NewOpenFGAClient(openFgaUrl string) (*OpenFgaClient, error) {
	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiUrl: openFgaUrl,
	})
	if err != nil {
		log.Fatalf("Erro ao inicializar o cliente OpenFGA: %v", err)
	}

	return fgaClient, nil
}

func checkStoreExists(ctx context.Context, client *OpenFgaClient) (*string, error) {
	response, err := client.ListStores(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't list stores: %v", err)
	}

	for _, store := range response.Stores {
		if store.Name == "Darwin" {
			return &store.Id, nil
		}
	}

	return nil, nil
}

func createStore(ctx context.Context, client *OpenFgaClient) error {
	_, err := client.CreateStore(ctx).Body(ClientCreateStoreRequest{
		Name: "Darwin",
	}).Execute()

	if err != nil {
		return fmt.Errorf("couldn't create store: %v", err)
	}

	return nil
}

func Initialize(ctx context.Context, client *OpenFgaClient) error {
	storeId, err := checkStoreExists(ctx, client)
	if err != nil {
		return fmt.Errorf("couldn't check if store exists: %v", err)
	}

	if storeId != nil {
		return fmt.Errorf("store already exists")
	}

	err = createStore(ctx, client)
	if err != nil {
		return err
	}

	return nil
}
