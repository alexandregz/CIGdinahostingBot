// implementacion de dir en sqlite
//
//	(Template design, aqui implementamos un dos providers)
package providers

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type DirSqlite struct {
	Dlf
}

// sem implementar, nom fai falha
func (d *DirSqlite) listFiles(path string) ([]string, error) {
	log.Fatal("Nom fai falha implementar")
	return nil, nil
}

// sem implementar, nom fai falha
func (d *DirSqlite) listDirectories(path string) ([]string, error) {
	log.Fatal("Nom fai falha implementar")
	return nil, nil
}

// devolve todas as rows, separando campos por pipes (|) e cada fila por um salto de linha
func (d *DirSqlite) getFile(taboa string, ficheiro string) ([]byte, error) {
	var comezo, fin, descripcion sql.NullString // "nullables"
	var tema string                             // NOT NULL

	// como é simple demostracion de como usar un proveedor de SQL, picamos todo aqui
	db, err := sql.Open("sqlite", ficheiro)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + taboa)
	if err != nil {
		panic(err)
	}

	var str string
	for rows.Next() {
		err = rows.Scan(&comezo, &fin, &tema, &descripcion)
		if err != nil {
			panic(err)
		}

		str += fmt.Sprintf("%s|%s|%s|%s\n", comezo.String, fin.String, tema, descripcion.String)
	}

	return []byte(str), nil
}

// sem implementar, nom fai falha
func (d *DirSqlite) exportAndDownloadFile(string, string, string) (string, error) {
	log.Fatal("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirSqlite) downloadFile(string, string) (string, error) {
	log.Fatal("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirSqlite) getSpreadsheetTab(string, string, string) ([][]interface{}, error) {
	log.Fatal("Nom fai falha implementar")
	return nil, nil
}
