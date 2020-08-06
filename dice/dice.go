package dice

import (
	"fmt"

	"github.com/NeapolitanIcecream/bot"
)

func hello(command *bot.Cmd) (msg string, err error) {
	var message = command.Message
	fmt.Println(message)
	// matched, err := regexp.MatchString(`foo.*`, "seafood")
	msg = fmt.Sprintf("Hello %s", command.User.RealName)
	return
}

func init() {
	fmt.Println("[Debug] Plugin \"dice\" loaded.")

	bot.RegisterCommand(
		"r",
		"Roll [n]d[m] dice. Example: \"!r 1d4+1d6+2\".",
		"",
		hello)
}
