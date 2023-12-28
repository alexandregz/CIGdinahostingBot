// menú de /axuda e /help
package menus

// devolve menu de /axuda|/help
func Axuda() string {
	return `/axuda|/help - <i>Este menú</i>
/asembleas - <i>Asembleas pasadas e a organizar</i>
/comite - <i>Temas do Comité de Empresa</i>
/comunicacions - <i>		Comunicacións por email comité/CIG/dinaworkers</i>
/css - <i>		Comité de Seguridade e Saúde</i>
/documentacion - <i>		Documentación: teletraballo, rexistros xornada, vacacións, convenios</i>
/igualdade - <i>Igualdade</i>
/lexislacion - <i>Documentación: teletraballo, PRL, Estatuto, convenios</i>
/plantilla - <i>		Lista traballadores/as</i>
/temas - <i>Temas en curso</i>`
}

// menú de admin, engádese a Axuda()
func AxudaAdmin() string {
	return Axuda() + `

<u>Admins:</u>
/canle		- <i>		Envía COMANDO a canle (por exemplo "/canle temas")</i>
/contasanuais		- <i>	Contas anuais presentadas pola empresa</i>
/horassindicais		<i>		Xestión horas</i>
/msg		- <i>		Envía mensagem (por exemplo "/canle msg hola mundo!")</i>
`
}
