package bot

import (
	"reflect"
	"testing"
)

// debe chamarse empregando o script que contem o token:
//
// alex@vosjod:~/Development/CIGdinahostingBot/src/bot(master)$ source ../../add_api_token_to_env.sh &&  go test -test.v
// === RUN   TestGetInstanceBot
// 2023/10/29 22:33:20 Endpoint: getMe, params: map[]
// 2023/10/29 22:33:20 Endpoint: getMe, response: {"ok":true,"result":{"id":xxxx,"is_bot":true,"first_name":"CIGdhBot","username":"CIGdhBot","can_join_groups":true,"can_read_all_group_messages":false,"supports_inline_queries":false}}
// --- PASS: TestGetInstanceBot (0.29s)
// PASS
// ok  	github.com/alexandregz/CIGdinahostingBot/src/bot	0.939s
func TestGetInstanceBot(t *testing.T) {

	botTG := GetInstanceBot()

	// da antes un panic(), pero a maiores aqui comprobamos se o tipo de botTG é un "struct"
	if reflect.TypeOf(botTG).Elem().Kind() != reflect.Struct {
		t.Fatal("botTG nom é de tipo struct!: " + reflect.TypeOf(botTG).Elem().Kind().String())
	}

	// um pouco de infor
	botTG.Debug = true
	botTG.GetMe()
}
