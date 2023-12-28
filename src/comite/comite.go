// comandos de /comite
package comite

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
	"github.com/alexandregz/CIGdinahostingBot/src/utils"
)

// menu de /comite
func menu() string {
	return `/comite actas
/comite acta DOC1			- <i>	Descarga acta de CSS</i>

/comite temas			- <i>	Temas en curso do Comité</i>
/comite regulamento		- 	<i>	Regulamento</i>`
}

// busca nun directorio pasado, copiado de /lexislacion
func menuListadoFicheros(directorio string) string {
	str := ""

	dirs, err := providers.ListFilesDirectory(directorio) // diretorio compartido via conta de servizo
	if err != nil {
		return "<i>Erro ao consultar " + directorio + ". " + err.Error() + "</i>"
	}
	for i := 0; i < len(dirs); i++ {
		str = str + " - " + dirs[i] + "\n"
	}

	return "<i>" + str + "</i>"
}

// descarga ficheiro desde path
func descargaFicheroPath(chatID int64, userID int64, command string, ficheiro string, diretorio string) {
	slog.Info(fmt.Sprintf("(user.ID %d) Descargando %s", userID, ficheiro))
	pathLocal, err := providers.GetFileOrExportToPDFLocal(ficheiro, diretorio)

	if err != nil {
		slog.Info("Erro ao descargar: " + err.Error())

		msg := "<i>Erro ao descargar " + ficheiro + " desde " + diretorio + ": "
		// a admins informamos do erro en concreto
		if user.IsAdmin(userID) {
			msg += ". " + err.Error()
		}
		msg += "</i>"
		senders.Send(msg, chatID, userID, command)
	} else {
		slog.Info(fmt.Sprintf("(user.ID %d) Ok. Enviando ficheiro %s", userID, pathLocal))
		senders.SendFile(chatID, pathLocal)
	}
}

// menu de /comite temas
func menuTemas() string {
	// buscamos spreadsheet "Temas en curso" no diretorio "Temas en curso"
	b, err := providers.GetFileGCPxlsx("Temas comite", "Temas comite", "Temas en curso")
	if err != nil {
		return "<i>Erro ao consultar Temas COMITE. " + err.Error() + "</i>"
	}

	// output, pode variar asi que nom controlamos, so fazemos output por debug
	str := ""
	for column, row := range b {
		if column == 0 {
			continue
		}

		// desc pode nom existir
		desc := ""
		if len(row) >= 4 {
			desc = strings.TrimSuffix(row[3].(string), "\n")
		}

		fechaComezo := row[0].(string)
		fechaFin := row[1].(string)
		tema := row[2].(string)

		str = str + " -"
		if len(fechaComezo) > 0 {
			str = str + fmt.Sprintf(" (Comezo: %s)", fechaComezo)
		}
		if len(fechaFin) > 0 {
			str = str + fmt.Sprintf(" (Fin: %s) ", fechaFin)
		}
		if len(tema) > 0 {
			str = str + fmt.Sprintf(" <b>%s</b> ", tema)
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

// envio de mensagens, parseando subcomandos se venhem, ver menu()
func HandleCommandComite(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "ACTAS" {
		senders.Send(menuListadoFicheros("Actas reuniones comité"), chatID, userID, command)
	} else if arg_principal == "ACTA" {
		descargaFicheroPath(chatID, userID, command, strings.Join(argumentos[1:], " "), "Actas reuniones comité")
	} else if arg_principal == "REGULAMENTO" {
		descargaFicheroPath(chatID, userID, command, "Regulamento_Comite_Dinahosting_novembro_2022", "regulamentos_comites_e_acta_constitucions")
	} else if arg_principal == "TEMAS" {
		senders.Send(menuTemas(), chatID, userID, command)
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
