package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendDiscordMessage(webhook, botName, msg string) error {
	data := map[string]interface{} {
		"content": msg,
	}

	if len(botName) > 0 {
		data["username"] = botName
	}

	req, _ := json.Marshal(data)

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(respData))

	return nil
}
