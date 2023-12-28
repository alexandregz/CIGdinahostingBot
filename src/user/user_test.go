package user

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	if ConfigFile == "" {
		ConfigFile = os.Getenv("FILE_CONFIG_USERS")
	}
	Config = ReadConfig()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCheckAdmins(t *testing.T) {
	firstAdmin := Config["admins"][0]

	if reflect.TypeOf(firstAdmin).Kind() != reflect.Int64 {
		t.Fatal("O primeiro admin devolto da config non é de tipo int64: " + reflect.TypeOf(firstAdmin).Kind().String())
	}
}

func TestCheckUsers(t *testing.T) {
	firstUser := Config["users"][0]

	if reflect.TypeOf(firstUser).Kind() != reflect.Int64 {
		t.Fatal("O primeiro user devolto da config non é de tipo int64: " + reflect.TypeOf(firstUser).Kind().String())
	}
}

func TestCheckSuperadmins(t *testing.T) {
	firstSuperadmin := Config["superadmins"][0]

	if reflect.TypeOf(firstSuperadmin).Kind() != reflect.Int64 {
		t.Fatal("O primeiro superadmin devolto da config non é de tipo int64: " + reflect.TypeOf(firstSuperadmin).Kind().String())
	}
}
