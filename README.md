# slack-notify

Simple Slack Notifications

### Usage

```
Usage:
  slack-notify [command]

Examples:
SLACK_TOKEN=xoxb-... SLACK_CHANNEL=[Channel-name] ./slack-notify

Available Commands:
  file        Send file
  help        Help about any command
  msg         Send messages

Flags:
  -h, --help   help for slack-notify

```

##### Flag `msg`
```
Usage:
  slack-notify msg [flags]

Flags:
  -b, --code-block    Send as code block
  -f, --file string   Send content of file as message.
  -h, --help          help for msg
  -t, --text string   Content of messages
```

##### Flag `file`
```
Usage:
  slack-notify file [flags]

Flags:
  -f, --file string    File path
  -h, --help           help for file
  -t, --title string   Title of message (Default is filepath)
```