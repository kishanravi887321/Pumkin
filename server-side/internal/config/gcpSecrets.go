package config

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func GetSecret(secretName, version string) (string, error) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		return "", fmt.Errorf("GCP_PROJECT_ID is not configured")
	}

	if secretName == "" || version == "" {
		return "", fmt.Errorf("secret name and version are required")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("create Secret Manager client: %w", err)
	}
	defer client.Close()

	name := fmt.Sprintf(
		"projects/%s/secrets/%s/versions/%s",
		projectID,
		secretName,
		version,
	)

	result, err := client.AccessSecretVersion(
		ctx,
		&secretmanagerpb.AccessSecretVersionRequest{
			Name: name,
		},
	)
	if err != nil {
		return "", fmt.Errorf("access secret %q: %w", secretName, err)
	}

	value := strings.TrimSpace(string(result.Payload.Data))
	if value == "" {
		return "", fmt.Errorf("secret %q is empty", secretName)
	}

	return value, nil
}
