// implementacion de dir para Nextcloud
//
//	(Template design, aqui implementamos un dos providers)
package providers

import (
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
)

// hard-coded aqui porque é un simple exemplo O:-)
var (
	url = "https://xx.xx.com:6443" // porto aberto no router para acceder desde fóra
	//url      = "192.168.xx.xx"
	username = "xx"
	password = "xx"
)

type DirNextcloud struct {
	Dlf
}

// sem implementar, nom fai falha
func (d *DirNextcloud) listFiles(path string) ([]string, error) {
	c, err := gonextcloud.NewClient(url)
	if err != nil {
		panic(err)
	}
	if err := c.Login(username, password); err != nil {
		panic(err)
	}
	defer c.Logout()

	var listaFicheiros []string

	folders, err := c.WebDav().ReadDir(".")
	if err != nil {
		panic("Unable to list files: " + err.Error())
	}
	for _, file := range folders {
		if !file.IsDir() {
			listaFicheiros = append(listaFicheiros, file.Name())
		}
	}

	return listaFicheiros, nil
}

// sem implementar, nom fai falha
func (d *DirNextcloud) listDirectories(path string) ([]string, error) {
	c, err := gonextcloud.NewClient(url)
	if err != nil {
		panic(err)
	}
	if err := c.Login(username, password); err != nil {
		panic(err)
	}
	defer c.Logout()

	var listaDirectorios []string

	folders, err := c.WebDav().ReadDir(".")
	if err != nil {
		panic("Unable to list files: " + err.Error())
	}
	for _, file := range folders {
		if file.IsDir() {
			listaDirectorios = append(listaDirectorios, file.Name())
		}
	}

	return listaDirectorios, nil
}

// devolve todas as rows, separando campos por pipes (|) e cada fila por um salto de linha
func (d *DirNextcloud) getFile(taboa string, ficheiro string) ([]byte, error) {
	panic("A implementar")
	return nil, nil
}

// sem implementar, nom fai falha
func (d *DirNextcloud) exportAndDownloadFile(string, string, string) (string, error) {
	panic("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirNextcloud) downloadFile(string, string) (string, error) {
	panic("Nom fai falha implementar")
	return "", nil
}

// sem implementar, nom fai falha
func (d *DirNextcloud) getSpreadsheetTab(string, string, string) ([][]interface{}, error) {
	panic("Nom fai falha implementar")
	return nil, nil
}
