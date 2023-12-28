package senders

import (
	"os"
	"strconv"
	"testing"

	"github.com/alexandregz/CIGdinahostingBot/src/canle"
	"github.com/alexandregz/CIGdinahostingBot/src/user"
)

// recolhido de config.yml
var userID int64

// chatID ven exportado em BOT_TG_CIGDH_CANLE_ID_DEFAULT
var chatID = canle.CanleChatID

func TestMain(m *testing.M) {
	if user.ConfigFile == "" {
		user.ConfigFile = os.Getenv("FILE_CONFIG_USERS")
	}
	user.Config = user.ReadConfig()

	userID = user.GetSuperAdmin()

	if chatID == 0 {
		chatID2, err := strconv.ParseInt(os.Getenv("BOT_TG_CIGDH_CANLE_ID_DEFAULT"), 10, 64)
		if err != nil {
			chatID = 0 // pasamos un valor por defecto, pero podemos forzalo
		} else {
			chatID = chatID2
		}
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

// hai que autenticarse, polo que necesita exportar o token:
//
// alex@vosjod:~/Development/CIGdinahostingBot/src/senders(master)$ source ../../add_api_token_to_env.sh && go test -test.v
// === RUN   TestSend
// --- PASS: TestSend (0.58s)
// PASS
// ok  	github.com/alexandregz/CIGdinahostingBot/src/senders	1.010s
func TestSend(t *testing.T) {
	// msg normal
	Send("test de Send()", userID, userID, "test")
}

func TestSendMsgHTMLChatID(t *testing.T) {
	_, err := SendMsgHTMLChatID(chatID, "test de SendMsgHTMLChatID()")

	if err != nil {
		t.Errorf("Erro enviando msg a canle: %s", err)
	}
}
