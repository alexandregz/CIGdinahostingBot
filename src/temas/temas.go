// comandos de /temas
package temas

import (
	"fmt"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/utils"
)

// menu de /temas
// func menu() string {
// 	return `/temas lista
// /temas engadir TEMA		- <i>	(ToDo) Engadir tema</i>
// /temas borrar TEMA 		- <i>	(ToDo) Borrar tema</i>`
// }

// menu de /temas lista
func menuTemas() string {
	// buscamos spreadsheet "Temas en curso" no diretorio "Temas en curso"
	b, err := providers.GetFileGCPxlsx("Temas en curso LISTADO", "Temas en curso", "Temas en curso")
	if err != nil {
		return "<i>Erro ao consultar Temas. " + err.Error() + "</i>"
	}

	// output, pode variar asi que nom controlamos, so fazemos output por debug
	str := ""
	for column, row := range b {
		if column == 0 {
			continue
		}

		fechaComezo := ""
		fechaFin := ""
		tema := ""
		desc := ""

		if len(row) > 0 {
			fechaComezo = row[0].(string)
		}
		if len(row) > 1 {
			fechaFin = row[1].(string)
		}
		if len(row) > 2 {
			tema = row[2].(string)
		}
		if len(row) > 3 {
			desc = strings.TrimSuffix(row[3].(string), "\n")
		}

		str = str + " -"
		if len(fechaComezo) > 0 {
			str = str + fmt.Sprintf(" (Comezo: %s)", fechaComezo)
		}
		if len(fechaFin) > 0 {
			str = str + fmt.Sprintf(" (Fin: %s) ", fechaFin)
		}
		if len(tema) > 0 {
			str = str + "<b>" + fmt.Sprintf(" %s ", tema) + "</b>"
		}
		if len(desc) > 0 {
			str = str + fmt.Sprintf(" ==> %s ", desc)
		}
		str = str + "\n"
	}

	// link a docu. Despois de creado o link elimino o seu valor para que nom se poida consumir incorretamente
	// providers.IdFile
	str = "<i>" + str + "</i>"

	str = str + "\nDocumento: " + utils.LinksGCP("spreadsheet", providers.IdFile)
	providers.IdFile = ""

	return str
}

// envio de mensagens, parseando subcomandos se venhem, ver Menu()
func HandleCommandTemas(chatID int64, userID int64, command string, args string) {
	// var arg_principal string
	// var argumentos []string
	// if args != "" {
	// 	argumentos = strings.Split(args, " ")
	// 	arg_principal = strings.ToUpper(argumentos[0])
	// }

	// polo de agora so hai /temas (nom "/temas lista", nem "/temas engadir TEMA", etc.)
	senders.Send(menuTemas(), chatID, userID, command)

	// if arg_principal == "" {
	// 	senders.Send(menu(), chatID, userID, command)
	// } else if arg_principal == "LISTA" {
	// 	senders.Send(menuTemas(), chatID, userID, command)
	// } else if arg_principal == "ENGADIR" {
	// 	// Todo
	// } else {
	// 	senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	// }
}
