// implementacion de dir en ficheiros locais
//
//	(Template design, aqui implementamos un dos providers)
package providers

import (
	"log"
	"os"
)

type DirLocal struct {
	Dlf
}

// lista SO directorios
func (d *DirLocal) listFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, e := range entries {
		if e.Type().IsRegular() {
			dirs = append(dirs, e.Name())
		}
	}

	return dirs, nil
}

// lista SO diretorios
func (d *DirLocal) listDirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, e := range entries {
		if e.Type().IsDir() {
			dirs = append(dirs, e.Name())
		}
	}

	return dirs, nil
}

// devolve o contido do ficheiro recibido
func (d *DirLocal) getFile(f string, path string) ([]byte, error) {
	content, error := os.ReadFile(path + "/" + f)
	if error != nil {
		panic(error)
	}

	return []byte(string(content)), nil
}

// sem implementar, nom fai falha
func (d *DirLocal) exportAndDownloadFile(string, string, string) (string, error) {
	log.Fatal("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirLocal) downloadFile(string, string) (string, error) {
	log.Fatal("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirLocal) getSpreadsheetTab(string, string, string) ([][]interface{}, error) {
	log.Fatal("Nom fai falha implementar")
	return nil, nil
}
