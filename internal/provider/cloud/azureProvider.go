package cloud

import (
	"errors"
	"fmt"

	"github.com/kubeopsskills/cloud-secret-resolvers/internal/provider"
	"github.com/kubeopsskills/cloud-secret-resolvers/internal/restapi"
)

type AzureProvider struct {
	Region     string
	SecretName string
	VaultURL   string

	API *restapi.AzureRestAPI

	accessToken *string
}

func (azProvider AzureProvider) InitialCloudSession() provider.CloudProvider {
	token, err := azProvider.API.GetAccessToken()
	if err != nil {
		fmt.Printf("Could not retrieve access token: %v", err)
	}
	azProvider.accessToken = token
	return azProvider
}

func (azProvider AzureProvider) RetrieveCredentials() (map[string]string, error) {
	data, err := azProvider.API.GetSecretValue(
		*azProvider.accessToken,
		azProvider.VaultURL,
		azProvider.SecretName,
	)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not retrieve any credentials: %v", err)
		return nil, errors.New(errorMessage)
	}
	return data, nil
}
