// CIGdhBot - Bot de Telegram para Comité de empresa de dinahosting
//
// Conecta a canle e resposta a certos comandos
// Alexandre Espinosa Menor - aemenor@gmail.com
//
// Uso: $ source ../add_api_token_to_env.sh && go run main.go [-debug Debug]
//
// 		alex@vosjod:~/Desktop/FP/FP PROXECTO/a20alexandreem/src(master)$ source ../add_api_token_to_env.sh && go run main.go -canle 'asminhasprobas_cigdhbot' -debug Debug
// 		INFO[2023-10-03T13:59:29+02:00] Start listening for updates. Press enter to stop
// 		INFO[2023-10-03T13:59:29+02:00] (user.ID xx) Nome wrote /igualdade
// 		INFO[2023-10-03T13:59:46+02:00] (user.ID xx) Nome wrote /igualdade@CIGdhBot
// 		INFO[2023-10-03T14:00:10+02:00] (user.ID xx) Nome wrote /axuda
// 		INFO[2023-10-03T14:00:18+02:00] (user.ID xx) Nome wrote /canle temas

// (2023-11-28): agora podense pasar todos os parametros necesarios por linha de comandos:
//
//  alex@vosjod:~/Desktop/FP/FP PROXECTO/a20alexandreem/src(argumentos_e_variables_entorno)$ go run main.go -token xx:xx -canle xx -chatID '-xx'
//			 -configFileUsers "$(pwd)/../config.yml" -gcpCredentialsFile "$(pwd)/../service-account.json"
//
// [-token 'xxx'] [-canle 'xxx''] [-chatID '-xxx'] [-configFileUsers /path/absoluto/config.yml] [-gcpCredentialsFile /path/absoluto/secure-account.json]

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strconv"

	"github.com/alexandregz/CIGdinahostingBot/src/bot"
	"github.com/alexandregz/CIGdinahostingBot/src/canle"
	"github.com/alexandregz/CIGdinahostingBot/src/handlers"
	"github.com/alexandregz/CIGdinahostingBot/src/providers"
	"github.com/alexandregz/CIGdinahostingBot/src/senders"
	"github.com/alexandregz/CIGdinahostingBot/src/user"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	debug string // tipo de log

	gcpCredentialsFile string // secure-account.json
	configFileUsers    string // config.yml
	token              string // token do bot
)

const PACKAGE_NAME = "CIGdinahostingBot"
const PACKAGE_VERSION = 1.0

func init() {
	// valor por defecto de variable de entorno (ver 'EXAMPLE add_api_token_to_env.sh')
	chatID, err := strconv.ParseInt(os.Getenv("BOT_TG_CIGDH_CANLE_ID_DEFAULT"), 10, 64)
	if err != nil {
		chatID = 0 // pasamos un valor por defecto, pero podemos forzalo
	}

	flag.StringVar(&canle.Canle, "canle", os.Getenv("BOT_TG_CIGDH_CANLE_PROBAS_DEFAULT"), "Canle a conectar, por exemplo 'probas_cigdhbot'")
	flag.StringVar(&debug, "debug", "", "'Debug', 'Info', 'Warning', 'Error' ou 'API' (debug de tgbotapi)")
	flag.Int64Var(&canle.CanleChatID, "chatID", chatID, "ChatID de canle a conectar, por exemplo '-1234567890'")

	flag.StringVar(&token, "token", "", "Token identificativo do bot")

	flag.StringVar(&configFileUsers, "configFileUsers", "", "Path ao ficheiro de config de usuarios")

	flag.StringVar(&gcpCredentialsFile, "gcpCredentialsFile", "", "Path do ficheiro de clave secreta de GCP")
}

func main() {
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// se os valores non venhem polos argumentos, cargamos valores do entorno
	if gcpCredentialsFile == "" {
		providers.FileCredentialsGCP = os.Getenv("CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP")
	} else {
		providers.FileCredentialsGCP = gcpCredentialsFile
	}

	if configFileUsers == "" {
		user.ConfigFile = os.Getenv("FILE_CONFIG_USERS")
	} else {
		user.ConfigFile = configFileUsers
	}

	if token == "" {
		bot.Token = os.Getenv("BOT_TG_CIGDH_APITOKEN")
	} else {
		bot.Token = token
	}
	// ----

	// cargo componhentes chave: users e gcp. Tinhaos en init() e cargase cando se importa o paquete, polo que se executaba antes de pedir args via CLI
	user.Config = user.ReadConfig()
	providers.CargaStructs()

	// opcions de log
	var programLevel = new(slog.LevelVar) // Info by default
	var addSource = false                 // AddSource

	if debug != "" {
		switch {
		case debug == "Debug":
			programLevel.Set(slog.LevelDebug)
			addSource = true
		case debug == "Info":
			addSource = true // INFO con AddSource
		case debug == "Warning":
			programLevel.Set(slog.LevelWarn)
			addSource = true
		case debug == "Error":
			programLevel.Set(slog.LevelError)
			addSource = true
		}
	}
	// -- log

	botTG := bot.GetInstanceBot()
	if debug == "API" {
		botTG.Debug = true
		programLevel.Set(slog.LevelDebug)
		addSource = true
	}

	// handler de slog. So amosa o nome do ficheiro porque nom debería haber duplicados!
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: addSource, Level: programLevel, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			s := a.Value.Any().(*slog.Source)
			s.File = path.Base(s.File)
		}
		return a
	}})
	slog.SetDefault(slog.New(h))
	slog.Info("Conectado")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	senders.SendMsgHTMLChatID(user.GetSuperAdmin(), fmt.Sprintf("[%s] Arrancando...", PACKAGE_NAME))

	// `updates` is a golang channel which receives telegram updates
	updates := botTG.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message updates
				continue
			}
			go handlers.HandleUpdate(update) // desplegamos nunha goroutine para atender a varios usuarios concurrentes
		}
	}
}
