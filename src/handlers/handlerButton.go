// handlerButton é o manexador dos callbacks, menús específicos de Telegram
package handlers

// handlerButton é o manexador dos callbacks, menús específicos de Telegram
// 		(exemplo https://github.com/alexandregz/CIGdinahostingBot/src/.iessanclemente.net/damd/a20alexandreem/-/raw/master/doc/img/menu_callback_query.png)

import (
	"github.com/alexandregz/CIGdinahostingBot/src/bot"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// tratamos as CallbackQuerys aqui (chega aqui ao pinchar no button)
func handleButton(query *tgbotapi.CallbackQuery) {
	botTG := bot.GetInstanceBot()

	callback := tgbotapi.NewCallback(query.ID, "")
	if _, err := botTG.Request(callback); err != nil {
		panic(err)
	}

	//enviamos mensagem para que introduzam datos
	senders.Send("Introduce "+query.Data+": ", query.From.ID, query.From.ID, "")
}
