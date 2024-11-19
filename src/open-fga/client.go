package open_fga

import (
	"context"
	"fmt"
	"log"

	. "github.com/openfga/go-sdk/client"
)

type Client struct {
	openFgaClient        *OpenFgaClient
	storeId              *string
	authorizationModelId *string
}

func getStoreId(ctx context.Context, client *OpenFgaClient) (*string, error) {
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

func getAuthorizationModelId(ctx context.Context, client *OpenFgaClient, storeId *string) (*string, error) {
	options := ClientReadLatestAuthorizationModelOptions{
		StoreId: storeId,
	}

	response, err := client.ReadLatestAuthorizationModel(ctx).Options(options).Execute()
	if err != nil {
		return nil, err
	}

	return &response.AuthorizationModel.Id, nil
}

func NewOpenFGAClient(openFgaUrl string) (*Client, error) {
	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiUrl: openFgaUrl,
	})
	if err != nil {
		log.Fatalf("couldn't initialize openfga client: %v", err)
	}

	storeId, err := getStoreId(context.Background(), fgaClient)
	if err != nil {
		return nil, err
	}

	authorizationModelId, err := getAuthorizationModelId(context.Background(), fgaClient, storeId)
	if err != nil {
		return nil, err
	}

	return &Client{openFgaClient: fgaClient, storeId: storeId, authorizationModelId: authorizationModelId}, nil
}

func (c *Client) Check(ctx context.Context, siteId string, userId string, action string) (*bool, error) {
	body := ClientCheckRequest{
		User:     fmt.Sprintf("user:%s", userId),
		Relation: "executor",
		Object:   fmt.Sprintf("action:%s", action),
		ContextualTuples: []ClientTupleKey{{
			User:     fmt.Sprintf("user:%s", userId),
			Relation: "member",
			Object:   fmt.Sprintf("site:%s", siteId),
		}},
	}

	options := ClientCheckOptions{
		AuthorizationModelId: c.authorizationModelId,
		StoreId:              c.storeId,
	}
	data, err := c.openFgaClient.Check(ctx).Body(body).Options(options).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't check authorization for user %s: %v", userId, err)
	}

	return data.Allowed, nil
}
