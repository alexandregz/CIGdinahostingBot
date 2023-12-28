// variables globais de canle
package canle

import (
	"os"
	"strconv"
)

var (
	Canle       = ""
	CanleChatID int64

	ComandoACanle = false // comando /canle. Envia umha vez, forzada, á canle e despois desativase (a false)
)

func init() {
	if CanleChatID == 0 {
		chatID, err := strconv.ParseInt(os.Getenv("BOT_TG_CIGDH_CANLE_ID_DEFAULT"), 10, 64)
		if err != nil {
			CanleChatID = 0 // pasamos un valor por defecto, pero podemos forzalo
		} else {
			CanleChatID = chatID
		}

	}
}
