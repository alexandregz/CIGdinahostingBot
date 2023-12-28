package temas

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
)

func TestMain(m *testing.M) {
	log.Println("Cargamos drivers de gcp()")
	providers.CargaStructs()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestMenuTemasGCP(t *testing.T) {

	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFile("temas.txt", "Temas en curso")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}
	if reflect.TypeOf(b).Elem().Kind() != reflect.Uint8 {
		t.Fatal("b nom é de tipo []byte!: " + reflect.TypeOf(b).Elem().Kind().String())
	}

	// output, pode variar asi que nom controlamos, so fazemos output por debug
	str := string(b[:])
	fmt.Printf("str: [%s]", str)
}

func TestMenuTemasGCPxlsx(t *testing.T) {

	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFileGCPxlsx("Temas en curso LISTADO", "Temas en curso", "Temas en curso")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}
	if reflect.TypeOf(b).Elem().Kind() != reflect.Slice {
		t.Fatal("b nom é de tipo []byte!: " + reflect.TypeOf(b).Elem().Kind().String())
	}
	fmt.Printf("b: [%#v]", b)
}

func TestMenuTemasSqlite(t *testing.T) {

	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFileExemploSQL("temas", "../providers/backend_sqlite.sqlite")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	if reflect.TypeOf(b).Elem().Kind() != reflect.Uint8 {
		t.Fatal("b nom é de tipo []byte!: " + reflect.TypeOf(b).Elem().Kind().String())
	}
	fmt.Printf("b: [%s]", b)
}

func TestMenuTemasLocal(t *testing.T) {

	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFileExemploLocal("backend_local.txt", "../providers/")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	if reflect.TypeOf(b).Elem().Kind() != reflect.Uint8 {
		t.Fatal("b nom é de tipo []byte!: " + reflect.TypeOf(b).Elem().Kind().String())
	}
	fmt.Printf("b: [%s]", b)
}
