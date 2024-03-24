# Simple Notifier

A simple CLI tool to send notification through webhooks

As of today, only Discord is implemented as that's what I'm using


## Configuration
The config file goes into `/etc/simple_notifier.yaml`

```yaml
locations:
    main_bot:
        type: discord
        bot_name: Backup bot
        webhook: WEBHOOK URL
```

## Usage

```
$ sn -l LOCATION_NAME -m "My message to be sent"
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
