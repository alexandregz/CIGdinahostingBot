// manexador de comandos
package handlers

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/asembleas"
	"github.com/alexandregz/CIGdinahostingBot/src/canle"
	"github.com/alexandregz/CIGdinahostingBot/src/comandos"
	"github.com/alexandregz/CIGdinahostingBot/src/comite"
	"github.com/alexandregz/CIGdinahostingBot/src/comunicacions"
	"github.com/alexandregz/CIGdinahostingBot/src/contasanuais"
	"github.com/alexandregz/CIGdinahostingBot/src/css"
	"github.com/alexandregz/CIGdinahostingBot/src/documentacion"
	"github.com/alexandregz/CIGdinahostingBot/src/igualdade"
	"github.com/alexandregz/CIGdinahostingBot/src/lexislacion"
	"github.com/alexandregz/CIGdinahostingBot/src/menus"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/temas"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// handler dos comandos disponhiveis
func handleCommand(chatID int64, userID int64, command string, args string) error {
	var err error

	// alex (2023-12-27): agora mesmo so deixamos a usuarios rexistrados executar comandos
	if !user.IsUser(userID) {
		senders.Send("Error ao executar comando: Non estás rexistrado como usuario", chatID, userID, command)

		slog.Info(fmt.Sprintf("Usuario sen rexistrar: [%d] (comando: /%s) (chatID: %d)", userID, command, chatID))
		return nil
	}

	// control de permisos antes de fazer calquera cousa
	if comandos.ComandoEParamDeAdmin(command, args) && !user.IsAdmin(userID) {
		senders.Send("Error ao executar comando: Non tes permisos", chatID, userID, command)

		slog.Info(fmt.Sprintf("Usuario NOM admin: [%d] (comando: /%s) (chatID: %d)", userID, command, chatID))
		return nil
	}

	switch strings.ToUpper(command) {
	case "AXUDA", "HELP":
		menuOutput := menus.Axuda()

		// para admins. E só por privado, nom pola canle!
		if user.IsAdmin(userID) && !canle.ComandoACanle && chatID == userID {
			menuOutput = menus.AxudaAdmin()
		}
		senders.Send(menuOutput, chatID, userID, command)

	case "TEMAS":
		temas.HandleCommandTemas(chatID, userID, command, args)

	case "IGUALDADE":
		igualdade.HandleCommandIgualdade(chatID, userID, command, args)

	case "CSS":
		css.HandleCommandCSS(chatID, userID, command, args)

	case "COMITE":
		comite.HandleCommandComite(chatID, userID, command, args)

	case "ASEMBLEAS":
		asembleas.HandleCommandAsembleas(chatID, userID, command, args)

	case "LEXISLACION":
		lexislacion.HandleCommandoConvenios(chatID, userID, command, args)

	case "DOCUMENTACION":
		documentacion.HandleCommandDocumentacion(chatID, userID, command, args)

	case "COMUNICACIONS":
		comunicacions.HandleCommandComunicacions(chatID, userID, command, args)

	case "CONTASANUAIS":
		contasanuais.HandleCommandContas(chatID, userID, command, args)

	// parsear um google calendar
	case "CALENDARIO":
		senders.Send("Calendario: ToDo", chatID, userID, command)

	// case "HORASSINDICAIS":
	// 	senders.Send(menus.HorasSindicais(), chatID, userID, command)

	case "MSG":
		senders.Send(args, chatID, userID, command)

	// case "PLANTILLA":
	// 	plantilla.HandleCommandPlantilla(chatID, userID, command, args)

	// informamos ao user por privado sempre
	default:
		senders.SendMsgHTMLChatID(userID, "Comando sen definir: "+command)
	}

	return err
}
