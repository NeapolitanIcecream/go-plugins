package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/NeapolitanIcecream/bot"
)

const errorMessage = "Error."

func diceFun(command *bot.Cmd) (msg string, err error) {
	message := strings.TrimLeft(command.Message, "r ")
	msgHead := fmt.Sprintf("%s roll %s: ", command.User.RealName, message)
	msgBody := ""
	diceResult := 0
	dices := strings.Split(message, "+")
	if len(dices) > 100 {
		msg = errorMessage
		return
	}

	for index, dice := range dices {
		numbers := strings.Split(strings.TrimSpace(dice), "d")
		submsg := ""

		switch len(numbers) {
		case 1:
			submsg = numbers[0]
			n, e := strconv.Atoi(numbers[0])
			if e != nil {
				msg = errorMessage
				return
			}
			diceResult += n

		case 2:
			n, e := strconv.Atoi(numbers[0])
			if e != nil || n > 50 {
				msg = errorMessage
				return
			}
			if n > 1 {
				submsg = strings.Join([]string{submsg, "("}, "")
			}
			m, e := strconv.Atoi(numbers[1])
			if e != nil || m > 1000 {
				msg = errorMessage
				return
			}
			result := rand.Int()%m + 1
			diceResult += result
			submsg = strings.Join([]string{submsg, fmt.Sprintf("%d", result)}, "")
			if n > 1 {
				for i := 1; i < n; i++ {
					result := rand.Int()%m + 1
					diceResult += result
					submsg = strings.Join([]string{submsg, fmt.Sprintf("%d", result)}, "+")
				}
				submsg = strings.Join([]string{submsg, ")"}, "")
			}

		default:
			msg = errorMessage
			return
		}

		if index > 0 {
			msgBody = strings.Join([]string{msgBody, submsg}, "+")
		} else {
			msgBody = strings.Join([]string{msgBody, submsg}, "")
		}
	}

	msgTail := fmt.Sprintf("=%v", diceResult)
	msg = strings.Join([]string{msgHead, msgBody, msgTail}, "")
	return
}

func init() {
	fmt.Println("[Debug] Plugin \"dice\" loaded.")

	bot.RegisterCommand(
		"r",
		"Roll [n]d[m] dice. Example: \"!r 1d4+1d6+2\".",
		"",
		diceFun)
}
