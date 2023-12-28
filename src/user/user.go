// parseo da config de usuarios
package user

import (
	"log/slog"
	"os"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

var (
	// Formato:
	//	admins:
	//	- 333
	//	- 444
	//	users:
	//	- 111
	//	- 2222
	// configFile = "/tmp/config.yml"
	ConfigFile string

	// mapa coa config: key => [valor1, valor2], key2...
	Config map[string][]int64
)

// leo ficheiro config.yml
func ReadConfig() map[string][]int64 {
	dat, err := os.ReadFile(ConfigFile)
	if err != nil {
		//panic("Erro lendo config.yml, path: " + ConfigFile)
		slog.Info("Erro lendo config.yml, path: " + ConfigFile)
	}

	data := make(map[string][]int64)

	err = yaml.Unmarshal([]byte(dat), &data)
	if err != nil {
		panic("erro en Unmarshall: " + err.Error())
	}
	return data
}

// devolve se o user é admin desde o arquivo de config.yml
func IsAdmin(userID int64) bool {
	return slices.Contains(Config["admins"][0:], userID) || IsSuperAdmin(userID)
}

// devolve se e superadmin
func IsSuperAdmin(userID int64) bool {
	return slices.Contains(Config["superadmins"][0:], userID)
}

// devolve se e superadmin
func IsUser(userID int64) bool {
	return slices.Contains(Config["users"][0:], userID)
}

// devolvemos o primeiro dos superadmins, abonda cum
func GetSuperAdmin() int64 {
	return Config["superadmins"][0]
}
