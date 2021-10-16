package service

import (
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

type VaultService interface {
	New() (*vault.Client, error)
	Read(path string) (*vault.Secret, error)
}

type VaultServiceImpl struct {
	Role   string
	Client *vault.Client
}

func (vaultService *VaultServiceImpl) New() (*vault.Client, error) {
	vaultConfig := vault.DefaultConfig()
	vaultClient, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Vault client: %w", err)
	}
	vaultJwt, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return nil, fmt.Errorf("unable to read file containing service account token: %w", err)
	}

	params := map[string]interface{}{
		"jwt":  string(vaultJwt),
		"role": vaultService.Role, // the name of the role in Vault that was created with this app's Kubernetes service account bound to it
	}
	// log in to Vault's Kubernetes auth method
	resp, err := vaultClient.Logical().Write("auth/kubernetes/login", params)
	if err != nil {
		return nil, fmt.Errorf("unable to log in with Kubernetes auth: %w", err)
	}
	if resp == nil || resp.Auth == nil || resp.Auth.ClientToken == "" {
		return nil, fmt.Errorf("login response did not return client token")
	}

	// now you will use the resulting Vault token for making all future calls to Vault
	vaultClient.SetToken(resp.Auth.ClientToken)
	vaultService.Client = vaultClient
	return vaultClient, nil
}

func (vaultService *VaultServiceImpl) Read(path string) (*vault.Secret, error) {
	// get secret from Vault
	return vaultService.Client.Logical().Read(path)
}
