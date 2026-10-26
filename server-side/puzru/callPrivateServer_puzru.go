package puzru

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

type PrivateServerResponse struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

func CallPrivateServer() (string, error) {
	url, err := config.GetPrivateServerURL()
	if err != nil {
		return "", err
	}

	// Use a timeout so the request cannot hang indefinitely.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("create private server request: %w", err)
	}

	authToken := os.Getenv("PRIVATE_SERVER_AUTH_TOKEN")
	if authToken == "" {
		return "", fmt.Errorf("PRIVATE_SERVER_AUTH_TOKEN is not configured")
	}
	req.Header.Set("X-Internal-Token", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("call private server: %w", err)
	}
	defer resp.Body.Close()

	// Check the HTTP status before decoding the response.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"private server returned status %d",
			resp.StatusCode,
		)
	}

	// Decode the JSON response into a Go struct.
	var result PrivateServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode private server response: %w", err)
	}

	return result.Key, nil
}
