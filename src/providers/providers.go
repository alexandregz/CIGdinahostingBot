// Loxica de consumo de GCP ou outros providers, para dar servizo aos temas
package providers

import (
	"log/slog"
	"os"
)

// devolve ficheiros do directorio compartido (sempre sem path, so o nome)
//
//	exemplo de cambio de proveedor (hard-coded pero podería programarse)
func ListFilesDirectory(d string) ([]string, error) {
	// provider GCP.
	dirGCP := &DirGCP{}
	p := Dlf{
		Dlf: dirGCP,
	}

	// // se empregamos outro provider, modificamos polo que toque tal que asi:
	// dirLocal := &DirLocal{}
	// p := Dlf{
	// 	Dlf: dirLocal,
	// }

	files, err := p.devolverListaFicheiros(d)

	if err != nil {
		return nil, err
	}
	return files, nil
}

// devolve []byte de ficheiro, por nome e path
func GetFile(f string, path string) ([]byte, error) {
	// provider GCP.
	dirGCP := &DirGCP{}

	p := Dlf{
		Dlf: dirGCP,
	}

	file, err := p.devolverFicheiro(f, path)

	if err != nil {
		return nil, err
	}
	// fmt.Printf("file: %#v", file)

	return file, nil
}

// devolve ruta local de ficheiro ou bem exportado e descargado ou so descargado, por nome e path
//
//	ficheiros de tipo .docx, .xlsx ou .pdf direitamente descargamse; ficheiros doutro tipo necesitam export previo
func GetFileOrExportToPDFLocal(f string, path string) (retFile string, err error) {
	slog.Debug("providers.GetFileOrExportToPDFLocal()")

	// control de posible panic() em p.Dlf.downloadFile()
	defer func() {
		err := recover()
		if err != nil {
			// provider GCP.
			dirGCP := &DirGCP{}

			p := Dlf{
				Dlf: dirGCP,
			}

			retFile, err = p.Dlf.exportAndDownloadFile(f, path, "application/pdf")
			if err != nil {
				retFile = ""
			}
		}
	}()

	// provider GCP.
	dirGCP := &DirGCP{}

	p := Dlf{
		Dlf: dirGCP,
	}

	// se da um panic() recuperamos o recover() e exportamos
	file, err := p.Dlf.downloadFile(f, path)
	if err != nil {
		return "", err
	}

	return file, nil
}

// devolve []byte de ficheiro, por nome e path de ficheiro directo de GCP (nom subido)
func GetFileGCPtxt(f string, path string) ([]byte, error) {
	// provider GCP.
	dirGCP := &DirGCP{}

	p := Dlf{
		Dlf: dirGCP,
	}
	// se da um panic() recuperamos o recover() e exportamos
	file, err := p.Dlf.exportAndDownloadFile(f, path, "text/plain")
	// file, err := p.Dlf.exportAndDownloadFile(f, path, "application/pdf")
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

// devolve []byte de spreadsheet, por nome e path de ficheiro directo de GCP (nom subido)
func GetFileGCPxlsx(f string, path string, tabAndRange string) ([][]interface{}, error) {
	// provider GCP.
	dirGCP := &DirGCP{}

	p := Dlf{
		Dlf: dirGCP,
	}
	// se da um panic() recuperamos o recover() e exportamos
	b, err := p.Dlf.getSpreadsheetTab(f, path, tabAndRange)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// EXEMPLO! devolve []byte de ficheiro, por nome e path. Simple copia e pega do GetFile(), cambiando o proveedor, para demostrar o uso do provider sqlite.go
func GetFileExemploSQL(taboa string, ficheiro string) ([]byte, error) {
	// provider sqlite.
	dirSqlite := &DirSqlite{}

	p := Dlf{
		Dlf: dirSqlite,
	}

	file, err := p.devolverFicheiro(taboa, ficheiro)

	if err != nil {
		return nil, err
	}
	// fmt.Printf("file: %#v", file)

	return file, nil
}

// EXEMPLO! devolve []byte de ficheiro, por nome e path
func GetFileExemploLocal(f string, path string) ([]byte, error) {
	// provider local
	dirLocal := &DirLocal{}
	p := Dlf{
		Dlf: dirLocal,
	}

	file, err := p.devolverFicheiro(f, path)

	if err != nil {
		return nil, err
	}
	// fmt.Printf("file: %#v", file)

	return file, nil
}

// EXEMPLO! devolve []byte de ficheiro, por nome e path
func GetFileExemploNextcloud(f string, path string) ([]byte, error) {
	// provider local
	dirNextcloud := &DirNextcloud{}
	p := Dlf{
		Dlf: dirNextcloud,
	}

	file, err := p.devolverFicheiro(f, path)

	if err != nil {
		return nil, err
	}
	// fmt.Printf("file: %#v", file)

	return file, nil
}

// EXEMPLO! lista diretorios de Nextcloud (da raiz do usuario)
func ListDirectoriesNextcloud(d string) ([]string, error) {
	// provider GCP.
	dirNextcloud := &DirNextcloud{}
	p := Dlf{
		Dlf: dirNextcloud,
	}

	files, err := p.Dlf.listDirectories(d)

	if err != nil {
		return nil, err
	}
	return files, nil
}

// EXEMPLO! lista ficheiros de Nextcloud
func ListFilesNextcloud(d string) ([]string, error) {
	// provider GCP.
	dirNextcloud := &DirNextcloud{}
	p := Dlf{
		Dlf: dirNextcloud,
	}

	files, err := p.Dlf.listFiles(d)

	if err != nil {
		return nil, err
	}
	return files, nil
}
