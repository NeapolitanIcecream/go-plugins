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

	for _, dice := range dices {
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
			if e != nil {
				msg = errorMessage
				return
			}
			if n > 1 {
				submsg = strings.Join([]string{submsg, "("}, "")
			}
			m, e := strconv.Atoi(numbers[1])
			if e != nil {
				msg = errorMessage
				return
			}
			result := rand.Int()%m + 1
			diceResult += result
			fmt.Printf("[Debug] n=%v, m=%v, result=%v\n", n, m, result)
			submsg = strings.Join([]string{submsg, fmt.Sprintf("%d", result)}, "")
			if n > 1 {
				for i := 1; i < n; i++ {
					result := rand.Int()%m + 1
					diceResult += result
					submsg = strings.Join([]string{submsg, string(result)}, "+")
				}
				submsg = strings.Join([]string{submsg, ")"}, "")
			}

		default:
			msg = errorMessage
			return
		}

		msgBody = strings.Join([]string{msgBody, submsg}, "+")
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
