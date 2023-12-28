// envio de mensaxes
package senders

import (
	"fmt"
	"log/slog"
	"path"

	"github.com/alexandregz/CIGdinahostingBot/src/bot"
	"github.com/alexandregz/CIGdinahostingBot/src/canle"
	"github.com/alexandregz/CIGdinahostingBot/src/comandos"
	"github.com/alexandregz/CIGdinahostingBot/src/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// discrimina se temos que enviar a canle ou a usuario
// os comandos vam sempre á mesma canle chatID (chatID é canle ou user), excepto:
//  1. que sexa um comando de admin
//  2. que venha fozado com /canle COMANDO
func Send(text string, chatID int64, userID int64, command string) {
	if canle.ComandoACanle {
		if canle.CanleChatID != 0 {
			SendMsgHTMLChatID(canle.CanleChatID, text)
		} else {
			SendMsgHTMLToChannel(canle.Canle, text)
		}
		canle.ComandoACanle = false // sempre desativamos
	} else {
		// se e comando admin ou vem de privado, a privado. Se nom enviamos a canle
		if comandos.ComandoDeAdmin(command) || chatID == userID {
			SendMsgHTMLChatID(userID, text)
		} else {
			if canle.CanleChatID != 0 {
				SendMsgHTMLChatID(canle.CanleChatID, text)
			} else {
				SendMsgHTMLToChannel(canle.Canle, text)
			}
		}
	}
}

// envía polo chatID (vale para canle tamém)
func SendMsgHTMLChatID(chatID int64, text string) (int, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML

	botTG := bot.GetInstanceBot()

	m, err := botTG.Send(msg)
	if err != nil {
		slog.Info(fmt.Sprintf("Erro enviando mensaxe a chat id %d: %s", chatID, err))
	}
	return m.MessageID, err
}

// envia a canle polo username (-canle '@xx XXXX XXXXX' funciona, aínda que daba erros antes )
func SendMsgToChannel(c string, text string) error {
	botTG := bot.GetInstanceBot()

	msg := tgbotapi.NewMessageToChannel(c, text)
	m, err := botTG.Send(msg)
	slog.Debug(fmt.Sprintf("enviado. chatid de channel: [%#v]", m.Chat.ChatConfig().ChatID))

	canle.CanleChatID = m.Chat.ChatConfig().ChatID

	if err != nil {
		panic(err)
	}

	return err
}

// envia a canle polo username desta
func SendMsgHTMLToChannel(c string, text string) error {
	msg := tgbotapi.NewMessageToChannel(c, text)
	msg.ParseMode = tgbotapi.ModeHTML

	botTG := bot.GetInstanceBot()
	m, err := botTG.Send(msg)

	canle.CanleChatID = m.Chat.ChatConfig().ChatID

	if err != nil {
		panic(err)
	}

	return err
}

// envia ficheiro. Por empregar chatID serve para canle
func SendFile(chatID int64, p string) error {
	botTG := bot.GetInstanceBot()

	// hack: nom damos enviado um ficheiro sem saber o chatID, por empregar NewDocument(chatID, filepath).
	// 	assi que, se estamos en /canle e o comando é de envio de ficheiro, enviamos antes umha mensagem á canle de que imos preparar o envío
	//	e desa maneira damos recolhido o chatID da canle, pois empregamos o NewMessageToChannel(canle, text)
	//
	//   DESCARTE se empregamos chatID (recomendado): agora pasamos o parametro "canleID", polo que pode que o tenhamos, si que nom fai falha
	//													enviar mensagem previa como hackeo
	if canle.ComandoACanle && canle.CanleChatID == 0 {
		SendMsgToChannel(canle.Canle, utils.SentenzasConformidade()+" Enviando ficheiro: "+path.Base(p))
	}

	if canle.ComandoACanle {
		chatID = canle.CanleChatID
		canle.ComandoACanle = false // sempre desativamos
	}

	_, err := botTG.Send(tgbotapi.NewDocument(chatID, tgbotapi.FilePath(p)))

	if err != nil {
		panic(err)
	}

	return err
}

// envio de mensagem con InlineKeyboard
func SendMessageWithReplyMarkup(chatID int64, text string, menu tgbotapi.InlineKeyboardMarkup) error {
	botTG := bot.GetInstanceBot()

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = menu

	if _, err := botTG.Send(msg); err != nil {
		panic(err)
	}
	return nil
}
