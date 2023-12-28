// handler de introducçom de datos
package handlers

import (
	"fmt"
	"log/slog"
)

// para replicar o comportamento do BotFather de Telegram, engadimos um modo de "introducción de datos": quando esteamos neste modo
// recolhemos os "nom comandos"
//
// Imolo
type IntroduccionDatos struct {
	Activada  bool              // se estamos en modo introducçom de datos ou nom
	Dato      string            // "path" do dato a introducir, despois engadese como key en Campos
	FromID    int64             // id do callback que nos chega
	MessageID int64             // id da mensagem que estamos a responder
	Campos    map[string]string // nome do campo a cubrir => valor
}

// activa o modo "Introduccion de datos"
func (v *IntroduccionDatos) ActivarIntroduccionDatos(dato string, chatID int64, messageID int64) {
	v.Activada = true
	v.Dato = dato
	v.FromID = chatID
	v.MessageID = messageID

	if v.Campos == nil {
		v.Campos = make(map[string]string)
	}

	slog.Debug(fmt.Sprintf("ActivarIntroduccionDatos v: [%#v]", v))
}

// desactiva o modo "Introduccion de datos", conserva Campos
func (v *IntroduccionDatos) DesactivarIntroduccionDatos() {
	v.Activada = false
	v.Dato = ""
	v.FromID = 0
	v.MessageID = 0

	slog.Debug(fmt.Sprintf("DesactivarIntroduccionDatos v: [%#v]", v))
}

// cancela o modo "Introduccion de datos"
func (v *IntroduccionDatos) CancelarIntroduccionDatos() {
	v.Activada = false
	v.Dato = ""
	v.FromID = 0
	v.MessageID = 0
	v.Campos = nil

	slog.Debug(fmt.Sprintf("CancelarIntroduccionDatos v: [%#v]", v))
}

// engade v.Dato co valor que se indique
func (v *IntroduccionDatos) EngadirCampo(val string) {
	if v.Campos == nil {
		v.Campos = make(map[string]string)
	}

	v.Campos[v.Dato] = val

	slog.Debug(fmt.Sprintf("EngadirCampo v: [%#v]", v))
}
