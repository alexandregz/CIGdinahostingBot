// manexador dos updates do channel de TG
package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// discrimina no update recibido entre messages e callbacks do menu
func HandleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		if update.Message.IsCommand() {
			handleMessage(update.Message)
		}

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
	}
}
