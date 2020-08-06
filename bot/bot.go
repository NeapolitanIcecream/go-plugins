package main

import (
	"os"

	_ "github.com/NeapolitanIcecream/go-plugins/dice"
	"github.com/go-chat-bot/bot/slack"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
