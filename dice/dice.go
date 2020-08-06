package dice

import (
	"fmt"

	"github.com/NeapolitanIcecream/bot"
)

func hello(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Hello %s", command.User.RealName)
	return
}

func init() {
	fmt.Println("[Debug] 0-0")

	bot.RegisterCommand(
		"hello",
		"Sends a 'Hello' message to you on the channel.",
		"",
		hello)
}
