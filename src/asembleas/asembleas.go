// comandos /asembleas
package asembleas

import (
	"fmt"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/utils"
)

// menu de /asembleas
func menu() string {
	return `/asembleas pasadas
/asembleas pasadas YYYY-MM-DD
/asembleas futuras`
}

// menu de /asembleas futuras|pasadas e /asembleas pasadas YYYY-MM-DD
//
//	se dia = "all", faise listado de todas as asembleas con todo
func menuAsembleas(tipo string, dia string) string {
	b, err := providers.GetFileGCPxlsx("Asembleas LISTADO", "Asembleas", tipo)
	if err != nil {
		return "<i>Erro ao consultar Asembleas. " + err.Error() + "</i>"
	}

	// output, pode variar asi que nom controlamos, so fazemos output por debug
	str := ""
	for column, row := range b {
		if column == 0 {
			continue
		}

		data := row[0].(string)
		if (dia != "" && data != dia) && dia != "all" {
			continue
		}
		desc := row[1].(string)
		asistentes := row[2].(string)

		// conclusions pode nom existir
		conclusions := ""
		if len(row) >= 4 {
			conclusions = row[3].(string)
		}

		str = str + " -"
		if len(data) > 0 {
			str = str + fmt.Sprintf("(Data: %s)", data)
		}
		if len(asistentes) > 0 {
			str = str + fmt.Sprintf(" (Asistentes: %s) ", asistentes)
		}

		if len(desc) > 0 {
			str = str + fmt.Sprintf(" %s ", desc)
		}

		// nom imos fazer output do resto se nom se pasa dia, so da data e asistentes
		if dia != "" || dia == "all" {
			if len(conclusions) > 0 {
				str = str + fmt.Sprintf(" \n >>> %s ", conclusions)
			}
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
func HandleCommandAsembleas(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "PASADAS" {
		if len(argumentos) == 1 {
			senders.Send(menuAsembleas("pasadas", ""), chatID, userID, command)
		} else {
			senders.Send(menuAsembleas("pasadas", argumentos[1]), chatID, userID, command)
		}
	} else if arg_principal == "FUTURAS" {
		senders.Send(menuAsembleas("futuras", ""), chatID, userID, command)
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
