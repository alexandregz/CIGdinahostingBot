// singleton de bot
package bot

import (
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var lock = &sync.Mutex{}

var (
	bot   *tgbotapi.BotAPI
	Token string
)

func init() {
	if Token == "" {
		Token = os.Getenv("BOT_TG_CIGDH_APITOKEN")
	}
}

func GetInstanceBot() *tgbotapi.BotAPI {
	if bot == nil {
		lock.Lock()
		defer lock.Unlock()

		var err error

		bot, err = tgbotapi.NewBotAPI(Token)
		if err != nil {
			panic(err)
		}

		// Set this to true to log all interactions with telegram servers
		bot.Debug = false
	}

	return bot
}
