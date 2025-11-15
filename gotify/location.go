package gotify

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/oxodao/simple-notifier/notification_service"
	"gopkg.in/yaml.v3"
)

type Location struct {
	BaseURL string `yaml:"base_url"`
	Token   string `yaml:"token"`
}

func (l Location) Send(msg notification_service.Message) error {
	values := url.Values{
		"message":  {msg.Content},
		"priority": {fmt.Sprintf("%v", msg.Priority)},
	}

	if len(msg.Title) > 0 {
		values.Add("title", msg.Title)
	}

	_, err := http.PostForm(l.BaseURL+"/message?token="+l.Token, values)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	notification_service.RegisterLocation(
		"gotify",
		func(settings []byte) (notification_service.Location, error) {
			var loc Location
			err := yaml.Unmarshal(settings, &loc)

			return loc, err
		},
	)
}
