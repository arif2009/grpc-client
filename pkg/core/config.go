package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var Config *config

// The JSON data stored in Secret Manager will follow this structure.
type config struct {
	ServiceSecret string `json:"service_secret"`
}

func Init(serviceName string) error {
	if os.Getenv("MODE") == "production" {
		return secretManagerInit(serviceName)
	}

	if err := localEnvInit(); err != nil {
		return err
	}

	return nil
}

// Get secrets from Secret Manager
func secretManagerInit(serviceName string) error {
	Config = &config{}
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", os.Getenv("PROJECT_ID"), serviceName),
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return err
	}

	if result.Payload != nil {
		if err := json.Unmarshal(result.Payload.Data, Config); err != nil {
			return err
		}
	}

	return nil
}

// For local dev, we'll either hard-code here or load from ENV
func localEnvInit() error {
	Config = &config{
		ServiceSecret: os.Getenv("SERVICE_SECRET"),
	}

	return nil
}
