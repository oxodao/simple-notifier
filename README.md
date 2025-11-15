# Simple Notifier

A super simple CLI tool to send notification through webhooks.

It aims at being as liteweight and easy to modify as possible.

Currently it features the following notification destinations:

- Discord
- Gotify

## Configuration

The config file goes into `/etc/simple_notifier.yaml`

```yaml
locations:
  main_bot:
    type: discord
    settings:
      bot_name: Backup bot
      webhook: WEBHOOK URL
```

## Usage

```
$ sn -l LOCATION_NAME [-t "My message title"] [-p PRIORITY] -m "My message to be sent"
```

Priority is a number between 0 and 10. 0 being silent notification, 5 the default means standard (vibrates / sound) and 10 being critical (shows the popup on Android for gotify).

Made to be used alongside [autorestic](https://github.com/cupcakearmy/autorestic), but it will probably fit any other workflow, idc do what you want with it.

## Bot types

### Discord

```yaml
locations:
  my_bot:
    type: "discord"
    settings:
      bot_name: "Backup bot"
      webhook: "Your channel webhook URL"
```

**Note**: For now, Discord does not supports the priority flag. This could be implemented as custom styling on the message or something but I did not do it, make a PR if you want.

### Gotify

```yaml
locations:
  my_bot:
    type: "gotify"
    settings:
      base_url: "https://gotify.example.com"
      token: "My Application Token"
```

## Adding a destination

In order to add a destination app, you will need to make a PR to this project adding your package containing the following:

```go
package my_service

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/oxodao/simple-notifier/notification_service"
	"gopkg.in/yaml.v3"
)

type Location struct {
    // Here is the settings your app will receive
}

func (l Location) Send(msg notification_service.Message) error {
	// Send the message to your service

	return nil
}

func init() {
	notification_service.RegisterLocation(
		"my_service", // The type it will be registered as
		func(settings []byte) (notification_service.Location, error) {
			var loc Location
			err := yaml.Unmarshal(settings, &loc)

			return loc, err
		},
	)
}
```

And then you simply need to import it in main.go:

```go
	// Loading all the available notification services
	_ "github.com/oxodao/simple-notifier/discord"
	_ "github.com/oxodao/simple-notifier/gotify"
	_ "github.com/oxodao/simple-notifier/my_service"
```

## License

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE

TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

0. You just DO WHAT THE FUCK YOU WANT TO.
