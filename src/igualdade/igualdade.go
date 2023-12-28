// comandos de /igualdade
package igualdade

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// menu de /igualdade
func menu() string {
	return `/igualdade actas
/igualdade acta YYYY-MM-DD		<i>	Descarga do día indicado
		(Se hai varias do mesmo día indica o nome do ficheiro aprox.)</i>

/igualdade plans		<i>	Plans de Igualdade de dinahosting</i>
/igualdade plan DOC1		<i>	Descarga dun PI</i>

/igualdade diagnose		<i>	Docs diagnose do Comité de Igualdade</i>
/igualdade diagnose DOC1		<i>	Descarga de doc</i>`
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

// envio de mensagens, parseando subcomandos se venhem, ver Menu()
func HandleCommandIgualdade(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "ACTAS" {
		senders.Send(menuListadoFicheros("Actas"), chatID, userID, command)
	} else if arg_principal == "ACTA" {
		if len(argumentos) != 2 {
			senders.Send("Non se recibiu correctamente a acta a enviar", chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "Actas")
		}
	} else if arg_principal == "PLANS" {
		senders.Send(menuListadoFicheros("PlansIgualdade"), chatID, userID, command)
	} else if arg_principal == "PLAN" {
		if len(argumentos) != 2 {
			senders.Send("Non se recibiu correctamente o plan a enviar", chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "PlansIgualdade")
		}
	} else if arg_principal == "DIAGNOSE" {
		if len(argumentos) == 1 {
			senders.Send(menuListadoFicheros("Diagnose"), chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "Diagnose")
		}
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
