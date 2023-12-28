// manexador de mensaxes
package handlers

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/canle"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text
	command := message.Command()
	args := message.CommandArguments()

	if user == nil {
		return
	}

	// Print to console
	slog.Info(fmt.Sprintf("(user.ID %d) %s wrote %s", user.ID, user.FirstName, text))

	if command == "canle" {
		res := strings.Split(args, " ")
		command = res[0]
		args = strings.Join(res[1:], " ")

		// envío único e forzado a canle, ver senders.Send()
		canle.ComandoACanle = true
	}

	var err error
	if command != "" {
		err = handleCommand(message.Chat.ID, user.ID, command, args)
	}

	if err != nil {
		slog.Info(fmt.Sprintf("An error occured: %s", err.Error()))
	}
}
