package comunicacions

import (
	"fmt"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/utils"
)

// menu de /asembleas
func menu() string {
	return `/comunicacions comite			<i>	Listado comunicacións como comité</i>
/comunicacions cig			<i>	Comunicacións como CIG</i>
/comunicacions listaamarela			<i>	Comunicacións dinaworkers</i>`
}

// menu de /comunicacions cig e /comunicacions comite
func menuComunicacionsSubmenu(tab string) string {
	b, err := providers.GetFileGCPxlsx("Comunicacions LISTADO", "Comunicacions", tab)
	if err != nil {
		return "<i>Erro ao consultar Comunicacions. " + err.Error() + "</i>"
	}

	// output, pode variar asi que nom controlamos, so fazemos output por debug
	str := ""
	for column, row := range b {
		if column == 0 {
			continue
		}

		data := row[0].(string)
		desc := row[1].(string)
		orixen := row[2].(string)
		// destinatarios := row[3].(string)
		raw := ""
		if len(row) >= 5 {
			raw = row[4].(string)
		}

		str = str + " -"
		if len(data) > 0 {
			str = str + fmt.Sprintf(" (%s)", data)
		}
		if len(orixen) > 0 {
			str = str + fmt.Sprintf(" (From: %s) ", orixen)
		}
		// if len(destinatarios) > 0 {
		// 	str = str + fmt.Sprintf(" (To: %s) ", destinatarios)
		// }
		if len(desc) > 0 {
			str = str + fmt.Sprintf(" %s ", desc)
		}
		if len(raw) > 0 {
			str = str + fmt.Sprintf(" %s ", raw)
		}
		str = str + "\n\n"
	}

	// link a docu. Despois de creado o link elimino o seu valor para que nom se poida consumir incorretamente
	// providers.IdFile
	str = "<i>" + str + "</i>"

	str = str + "\nDocumento: " + utils.LinksGCP("spreadsheet", providers.IdFile)
	providers.IdFile = ""

	return str
}

// envio de mensagens, parseando subcomandos se venhem, ver Menu()
func HandleCommandComunicacions(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "COMITE" {
		senders.Send(menuComunicacionsSubmenu("comite"), chatID, userID, command)
	} else if arg_principal == "CIG" {
		senders.Send(menuComunicacionsSubmenu("cig"), chatID, userID, command)
	} else if arg_principal == "LISTAAMARELA" {
		senders.Send(menuComunicacionsSubmenu("dinaworkers"), chatID, userID, command)
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
