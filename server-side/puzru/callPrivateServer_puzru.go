package puzru

import (
	"fmt"
	"net/http"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

func CallPrivateServer() (string, error) {
	url, err := config.GetPrivateServerURL()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to call private server: %d", resp.StatusCode)
	}

	return "Private server called successfully", nil
}
