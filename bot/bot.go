package main

import (
	"fmt"
	"io/ioutil"

	"github.com/NeapolitanIcecream/bot/slack"
	_ "github.com/NeapolitanIcecream/go-plugins/dice"
)

func main() {
	token, err := ioutil.ReadFile("../secrets/SLACK_TOKEN")
	if err != nil {
		fmt.Print(err)
	}
	slack.Run(string(token))
}
