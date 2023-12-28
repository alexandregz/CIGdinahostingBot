package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/alexandregz/CIGdinahostingBot/src/providers"
)

func TestMain(m *testing.M) {
	log.Println("Cargamos drivers de gcp()")
	providers.CargaStructs()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSentenzasConformidade(t *testing.T) {

	for i := 1; i <= 5; i++ {
		fmt.Println(SentenzasConformidade())
	}
}

func TestSentenzasDisconformidade(t *testing.T) {
	for i := 1; i <= 5; i++ {
		fmt.Println(SentenzasDisconformidade())
	}
}

// comprobamos que o link ao excel de Temas se devolve corretamente e coincide co anotado, asi como que e um link
func TestLinksGCP(t *testing.T) {

	linkTemas := LinksGCP("spreadsheet", "1VYILuqtjIoEvRl4XbEkXPuJcigfGPUpKKFif2ia-5Lg")

	_, err := providers.GetFileGCPxlsx("Temas en curso LISTADO", "Temas en curso", "Temas en curso")
	if err != nil {
		t.Fatal("Erro ao consultar Temas: " + err.Error())
	}

	linkTemasBusqueda := LinksGCP("spreadsheet", providers.IdFile)

	if linkTemas != linkTemasBusqueda {
		t.Fatalf("Os links de TEMAS non coinciden: [%s] - [%s]", linkTemas, linkTemasBusqueda)
	}

	//fmt.Printf("Links de TEMAS coinciden: [%s] - [%s]", linkTemas, linkTemasBusqueda)

	f, err := providers.GetFileOrExportToPDFLocal("Temas en curso LISTADO", "Temas en curso")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	// fmt.Println("f: " + f)

	// ---------------
	contentTypeExpected := "PDF document, version 1\\.4"

	contentType, err := GetFileContentType(f)
	if err != nil {
		t.Fatal("Erro ao consultar GetFileContentType(f): " + err.Error())
	}
	fmt.Println("Baixado xlsx de Temas e exportado a PDF. ContentType: " + contentType)

	if match, err := regexp.MatchString(contentTypeExpected, contentType); match == false || err != nil {
		t.Fatalf("Erro ao comprobar contentType . Recibido: [%s]. Agardado: [%s]. Erro: [%s]", contentType, contentTypeExpected, err)
	}
}

// comprobamos o filetype dun .txt
func TestGetFileContentType(t *testing.T) {
	f := "../providers/backend_local.txt"
	f2 := "../providers/backend_sqlite.sqlite"

	// en mac o file devolve "UTF-8 Unicode text", en Debian (imagem golang:latest de github.com/alexandregz/CIGdinahostingBot/src/-ci) devolve "Unicode text, UTF-8 text", asi que comparo so UTF-8
	contentExpected := "UTF-8" // "UTF-8 Unicode text"
	contentExpected2 := "SQLite 3.x database, last written using SQLite version 3034001"

	contentType, err := GetFileContentType(f)
	if err != nil {
		t.Fatal("Erro ao consultar GetFileContentType(f): " + err.Error())
	}
	contentType2, err2 := GetFileContentType(f2)
	if err2 != nil {
		t.Fatal("Erro ao consultar GetFileContentType(f2): " + err2.Error())
	}

	if match, err := regexp.MatchString(contentExpected, contentType); match == false || err != nil {
		t.Fatalf("Erro ao comprobar contentType . Recibido: [%s]. Agardado: [%s]. Erro: [%s]", contentType, contentExpected, err)
	}

	if match, err := regexp.MatchString(contentExpected2, contentType2); match == false || err != nil {
		t.Fatalf("Erro ao comprobar contentType . Recibido: [%s]. Agardado: [%s]. Erro: [%s]", contentType2, contentExpected2, err)
	}

	// fmt.Printf("contentType de %s OK. Agardado: [%s]. Devolto: [%s]\n", f, contentExpected, contentType)
	// fmt.Printf("contentType de %s OK. Agardado: [%s]. Devolto: [%s]\n", f2, contentExpected2, contentType2)
}
