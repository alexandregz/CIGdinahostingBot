// comandos de /lexislacion
package lexislacion

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// menu de /documentacion
func menu() string {
	return `/lexislacion - <i>Este menú</i>

/lexislacion conveniosdh - <i>Convenios dinahosting</i>
/lexislacion conveniosdh YYYY - <i>Descarga</i>

/lexislacion PRL - <i>Lexislacion Prevención Riscos Laborais</i>
/lexislacion PRL FICHERO - <i>Descarga</i>

/lexislacion ET - <i>Estatuto Traballadores</i>
/lexislacion ET FICHERO - <i>Descarga</i>

/lexislacion convenios - <i>Lexislacion Convenios</i>
/lexislacion convenios FICHERO - <i>Descarga</i>

/lexislacion teletraballo - <i>Lexislacion teletraballo</i>
/lexislacion teletraballo FICHERO - <i>Descarga</i>`
}

// busca nun directorio pasado, para non repetir 5 veces o mesmo codigo de listado dun directorio dado
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
func HandleCommandoConvenios(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "CONVENIOSDH" {
		if len(argumentos) > 1 {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "Convenios dinahosting")
		} else {
			senders.Send(menuListadoFicheros("Convenios dinahosting"), chatID, userID, command)
		}
	} else if arg_principal == "PRL" {
		if len(argumentos) > 1 {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "PRL")
		} else {
			senders.Send(menuListadoFicheros("PRL"), chatID, userID, command)
		}
	} else if arg_principal == "ET" {
		if len(argumentos) > 1 {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "Estatuto Traballadores")
		} else {
			senders.Send(menuListadoFicheros("Estatuto Traballadores"), chatID, userID, command)
		}
	} else if arg_principal == "CONVENIOS" {
		if len(argumentos) > 1 {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "Convenios")
		} else {
			senders.Send(menuListadoFicheros("Convenios"), chatID, userID, command)
		}
	} else if arg_principal == "TELETRABALLO" {
		if len(argumentos) > 1 {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "teletraballo")
		} else {
			senders.Send(menuListadoFicheros("teletraballo"), chatID, userID, command)
		}
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
