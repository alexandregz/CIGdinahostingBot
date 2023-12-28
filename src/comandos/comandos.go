// definición de comandos de admin
package comandos

import (
	"strings"

	"golang.org/x/exp/slices"
)

var (
	// comandos de Admin: a resposta so vai a privado e a admins
	//	para desactivar comandos "enteiros" (incluidos todos os subcomandos), engado * ao comando, ex: "msg *"
	comandosAdmin = [...]string{"msg *", "igualdade *", "temas privados", "css *", "comite *", "asembleas *", "documentacion *", "calendario *", "plantilla *",
		"comunicacions *", "horassindicais *", "canle *", "private *", "contasanuais *"}
)

// comproba se o comando + args ou "comando *" son de uso exclusivo de admins
func ComandoEParamDeAdmin(comando string, args string) bool {
	var arg_principal string
	var argumentos []string

	if args != "" {
		argumentos = strings.Split(args, " ")
		arg_principal = argumentos[0]
	}
	cmdAndArgPrincipal := comando + " " + arg_principal

	// desativado todo o comando
	cmdAndSubcommands := comando + " *"

	return (slices.Contains(comandosAdmin[0:], cmdAndArgPrincipal) || slices.Contains(comandosAdmin[0:], cmdAndSubcommands))
}

// comproba se o comando e de admin (ainda que nom coincida o arg principal)
func ComandoDeAdmin(comando string) bool {
	for _, v := range comandosAdmin {
		c := strings.Split(v, " ")

		if c[0] == comando {
			return true
		}
	}

	return false
}

// para test de senders
func GetComandoDeAdmin() string {
	return comandosAdmin[0]
}
