//
// Loxica de consumicion de GCP ou outros providers, consumidos polos adapters de cada tema (css, igualdade, comite, etc.)
//

package providers

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/alexandregz/CIGdinahostingBot/src/utils"
)

func TestMain(m *testing.M) {
	log.Println("Cargamos drivers de gcp()")
	CargaStructs()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetFileOrExportToPDFLocal(t *testing.T) {
	contentTypeExpected := "PDF document, version 1\\.5"

	// o ficheiro é un .docx de Google Cloud. Vaise baixar exportándoo a PDF
	f, err := GetFileOrExportToPDFLocal("2021-01-13", "Actas")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	fmt.Printf("Descargado: [%s]\n", f)

	contentType, err := utils.GetFileContentType(f)
	if err != nil {
		t.Fatal("Erro ao consultar GetFileContentType(f): " + err.Error())
	}

	// emprego regexp porque en golang:latest ao ser debian devolvía "PDF document, version 1.5, 2 pages (zip deflate encoded)" e en mac "PDF document, version 1.5"
	if match, err := regexp.MatchString(contentTypeExpected, contentType); match == false || err != nil {
		t.Fatalf("Erro ao comprobar contentType . Recibido: [%s]. Agardado: [%s]. Erro: [%s]", contentType, contentTypeExpected, err)
	}

	// comprobamos que un ficheiro false pete; se non dá erro, lanzamos un Fatal()
	_, err2 := GetFile("2021-01-13", "Actas FAKE")
	if err2 == nil {
		log.Fatal("Erro2: Non deu erro ao buscar un ficheiro falso!")
	}

}

// GetExportedFileToTXTLocal == descarga ficheiro exportado a .txt e devolve path ao ficheiro descargado (se está descargado nom o volve descargar)
func TestGetFileOrExportToPDFLocal2(t *testing.T) {

	// este ficheiro necesita export!
	// em os.TempDir() aparecem o ficheiro e o .pdf exportado (polo menos em mac)
	//   -rw-r--r--     1 alex  staff   47379 Nov  5 18:43 acta 26-01-23.pdf
	//   -rw-r--r--     1 alex  staff       0 Nov  5 18:44 acta 26-01-23
	f, err := GetFileOrExportToPDFLocal("acta 26-01-23", "Actas reuniones comité")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}
	pathAbsoluto := os.TempDir() + string(os.PathSeparator) + "acta 26-01-23.pdf"
	if pathAbsoluto != f {
		t.Errorf("Ficheiro descargado [%s] distinto ao agardado: [%s]", f, pathAbsoluto)
	}
	fmt.Printf("ficheiro1 em local (Exportado previamente): %s\n", f)

	// este ficheiro baixase direitamente
	f2, err2 := GetFileOrExportToPDFLocal("2021-01-13", "Actas")
	if err2 != nil {
		log.Fatal("Erro: " + err2.Error())
	}
	pathAbsoluto2 := os.TempDir() + string(os.PathSeparator) + "2021-01-13 ACTA DE LA COMISIÓN DE IGUALDAD.pdf"
	if pathAbsoluto2 != f2 {
		t.Errorf("Ficheiro descargado [%s] distinto ao agardado: [%s]", f2, pathAbsoluto2)
	}
	fmt.Printf("ficheiro2 em local (so descarga): %s\n", f2)

}

// Test de provider Nextcloud.
func TestGetListDirectoriesNextcloud(t *testing.T) {
	folders, err := ListDirectoriesNextcloud(".")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	if len(folders) != 3 {
		t.Errorf("O tamanho do array de directorios é distinto de 3: len = %d", len(folders))
	}

	if folders[0] != "Documents" {
		t.Errorf("O primeiro directorio non é Documents: %s", folders[0])
	}

	// for n, dir := range folders {
	// 	fmt.Printf("%d - %s\n", n, dir)
	// }
}

// Test de provider Nextcloud.
func TestGetListFilesNextcloud(t *testing.T) {
	files, err := ListFilesNextcloud("./Documents")
	if err != nil {
		log.Fatal("Erro: " + err.Error())
	}

	if len(files) != 6 {
		t.Errorf("O tamanho do array de directorios é distinto de 6: len = %d", len(files))
	}

	if files[0] != "Nextcloud Manual.pdf" {
		t.Errorf("O primeiro directorio non é Nextcloud Manual.pdf: %s", files[0])
	}

	// for n, file := range files {
	// 	fmt.Printf("%d - %s\n", n, file)
	// }
}
