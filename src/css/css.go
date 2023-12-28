package css

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// menu de /css
func menu() string {
	return `/css actas
/css acta YYYY-MM-DD

/css datas 				<i>	(ToDo) Datas de novas reunións e historico de velhas</i>
/css data YYYY-MM-DD "nova reunión" 				<i>	(ToDo) Engadir nova reunión / Modificar resumo de reunión</i>
/css data YYYY-MM-DD 				<i>	(ToDo) Resumo de reunión pasada</i>

/css documentacion
/css documentacion DOC1				<i>	Descarga</i>`
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
func HandleCommandCSS(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "ACTAS" {
		senders.Send(menuListadoFicheros("CSS actas"), chatID, userID, command)
	} else if arg_principal == "ACTA" {
		if len(argumentos) != 2 {
			senders.Send("Non se recibiu correctamente a acta a enviar", chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "CSS actas")
		}
	} else if arg_principal == "DOCUMENTACION" {
		if len(argumentos) == 1 {
			senders.Send(menuListadoFicheros("prevençóm delegados delegadas"), chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "prevençóm delegados delegadas")
		}

	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
