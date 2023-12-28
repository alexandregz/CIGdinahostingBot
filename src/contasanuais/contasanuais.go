package contasanuais

import (
	"strings"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// menu de /contasanuais
func menu() string {
	return `/contasanuais listado			<i>	Listado</i>
/contasanuais descarga DOC1			<i>	Descarga ficheiro de contas indicado (nome aprox.)</i>`
}

// menu de /contasanuais listado
func menuContas() string {
	str := ""

	dirs, err := providers.ListFilesDirectory("contas_anuais") // diretorio compartido via conta de servizo
	if err != nil {
		return "<i>Erro ao consultar Contas anuais. " + err.Error() + "</i>"
	}
	for i := 0; i < len(dirs); i++ {
		str = str + dirs[i] + "\n"
	}

	return "<i>" + str + "</i>"
}

// envio de mensagens, parseando subcomandos se venhem, ver Menu()
func HandleCommandContas(chatID int64, userID int64, command string, args string) {
	var arg_principal string
	var argumentos []string
	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = strings.ToUpper(argumentos[0])
	}

	if arg_principal == "" {
		senders.Send(menu(), chatID, userID, command)
	} else if arg_principal == "LISTADO" {
		senders.Send(menuContas(), chatID, userID, command)
	} else if arg_principal == "DESCARGA" {
		pathLocalActa, err := providers.GetFileOrExportToPDFLocal(strings.Join(argumentos[1:], " "), "contas_anuais")

		if err != nil {
			msg := "<i>Erro ao descargar ficheiro de Contas"
			if user.IsAdmin(userID) {
				msg += ". " + err.Error()
			}
			msg += "</i>"
			senders.Send(msg, chatID, userID, command)
		} else {
			senders.SendFile(chatID, pathLocalActa)
		}
	} else {
		senders.Send("Erro en subcomandos ou quizais non implementados?: "+args, chatID, userID, command)
	}
}
