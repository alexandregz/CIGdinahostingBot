// Template de navegacion de diretorios.
//
//	Para submenús de primeira orde
//	(Template pattern design, definimos aqui o Template)
package providers

type DirListFiles interface {
	listDirectories(path string) ([]string, error)                                           // lista directorios colgando do diretorio que se lhe pase (nom recursivo)
	listFiles(path string) ([]string, error)                                                 // lista ficheiros do diretorio pasado
	getFile(name string, path string) ([]byte, error)                                        // devolve []byte do ficheiro pasado, por (nome, carpeta da que colga). O nome e parte do nome do ficheiro, nom enteiro (em GCP)
	exportAndDownloadFile(name string, path string, mimetype string) (string, error)         // para GCP: devolve ruta absoluta dum ficheiro, exportandoo previamente antes da descarga
	downloadFile(name string, path string) (string, error)                                   // para GCP: descarga e devolve ruta absoluta do ficheiro temporal descargado
	getSpreadsheetTab(name string, path string, tabAndRange string) ([][]interface{}, error) // lectura dumha spreadsheet. tabAndRange, formato de sheets (tab!A1:E10)
}

type Dlf struct {
	Dlf DirListFiles
}

// devolve lista de subdirectorios
func (d *Dlf) devolverListaDirectorios(path string) ([]string, error) {

	dirs, err := d.Dlf.listDirectories(path)

	if err != nil {
		return nil, err
	}
	return dirs, nil

}

func (d *Dlf) devolverListaFicheiros(path string) ([]string, error) {

	files, err := d.Dlf.listFiles(path)

	if err != nil {
		return nil, err
	}
	return files, nil

}

// devolve ficheiro para empregar em senders.SendFile()
//
//	nome é o identificador do ficheiro (por exemplo YYYY-MM-DD no caso de actas) e tipo é o tipo de ficheiro,
//	normalmente o subcomando (por exemplo "acta" no caso de actas de IG)
//	pois recibese o comando "/igualdade acta YYYY-MM-DD" para enviar o ficheiro, polo que sería "nome=2022-01-01, tipo=acta"
func (d *Dlf) devolverFicheiro(nome string, tipo string) ([]byte, error) {

	files, err := d.Dlf.getFile(nome, tipo)

	if err != nil {
		return nil, err
	}
	return files, nil

}

// devolve []byte com datos de spreadsheet de GCP
func (d *Dlf) getSpreadsheetTab(nome string, tipo string, tabAndRange string) ([][]interface{}, error) {

	files, err := d.Dlf.getSpreadsheetTab(nome, tipo, tabAndRange)

	if err != nil {
		return nil, err
	}
	return files, nil

}
