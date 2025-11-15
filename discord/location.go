package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/oxodao/simple-notifier/notification_service"
	"gopkg.in/yaml.v3"
)

type Location struct {
	Webhook string `yaml:"webhook"`
	BotName string `yaml:"bot_name"`
}

func (l Location) Send(m notification_service.Message) error {
	data := map[string]any{}

	if len(l.BotName) > 0 {
		data["username"] = l.BotName
	}

	markdown := ""

	if len(m.Title) > 0 {
		markdown += "# **" + m.Title + "**\n\n"
	}

	markdown += m.Content

	data["content"] = markdown

	req, _ := json.Marshal(data)

	resp, err := http.Post(l.Webhook, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respStr, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.New("error sending a message to discord, but failed to read it: " + err.Error())
		}

		return errors.New("error sending a message to discord: " + string(respStr))
	}

	return nil
}

// Init is a special golang function that is called on package load
// thus we only need to import the discord package in the config
func init() {
	notification_service.RegisterLocation(
		"discord",
		func(settings []byte) (notification_service.Location, error) {
			var loc Location
			err := yaml.Unmarshal(settings, &loc)

			return loc, err
		},
	)
}
