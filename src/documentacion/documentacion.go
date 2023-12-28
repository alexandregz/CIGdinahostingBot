package documentacion

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
	return `/documentacion teletraballo			<i>	(ToDo) Documentación de teletraballo</i>

/documentacion rexistros_xornada
/documentacion rexistros_xornada YYYY-MM			<i>	Rexistro de xornada do mes indicado</i>

/documentacion vacacions
/documentacion vacacions YYYY			<i>	Descarga do ano indicado</i>

/documentacion copiasbasicas
/documentacion copiasbasicas YYYY-MM-DD			<i>	Descarga
	(Se hai varias do mesmo día indica o nome do ficheiro aprox.)</i>

/documentacion convenios			<i>	Convenios dinahosting</i>
/documentacion convenio YYYY			<i>	Descarga</i>

/documentacion calendarios			<i>Calendarios laborais Stgo/Madrid</i>
/documentacion calendario YYYY			<i>Descarga</i>

/documentacion rnt			<i>Ficheiros RNT/RLC disponhiveis</i>
/documentacion rnt YYYY			<i>	Descarga</i>`
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
func HandleCommandDocumentacion(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)

	} else if arg_principal == "CONVENIOS" {
		senders.Send(menuListadoFicheros("convenios"), chatID, userID, command)
	} else if arg_principal == "CONVENIO" {
		descargaFicheroPath(chatID, userID, command, argumentos[1], "convenios")

	} else if arg_principal == "CALENDARIOS" {
		senders.Send(menuListadoFicheros("calendarios_laborais"), chatID, userID, command)
	} else if arg_principal == "CALENDARIO" {
		descargaFicheroPath(chatID, userID, command, argumentos[1], "calendarios_laborais")

	} else if arg_principal == "COPIASBASICAS" {
		if len(argumentos) == 1 {
			senders.Send(menuListadoFicheros("copias_basicas"), chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "copias_basicas")
		}
	} else if arg_principal == "RNT" {
		if len(argumentos) == 1 {
			senders.Send(menuListadoFicheros("rnt_rlt"), chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "rnt_rlt")
		}
	} else if arg_principal == "VACACIONS" {
		if len(argumentos) == 1 {
			senders.Send(menuListadoFicheros("vacacions"), chatID, userID, command)
		} else {
			descargaFicheroPath(chatID, userID, command, argumentos[1], "vacacions")
		}
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
