// Providers empregados no backend da app
//
// implementacion de dir en GCP
//
//	(Template design, aqui implementamos un dos providers, o que tira contra Google Drive)
package providers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type DirGCP struct {
	Dlf
}

// struct Service que conecta com GCP
var cal *calendar.Service
var srv *drive.Service
var srvSheets *sheets.Service

// ID do file tratado, para poder crear links aos documentos umha vez consumido o doc
var IdFile string

var FileCredentialsGCP string

// aqui implementamos, nom como em dirFiles
func CargaStructs() {
	var err error

	if FileCredentialsGCP == "" {
		FileCredentialsGCP = os.Getenv("CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP")
	}
	//FileCredentialsGCP, _ = filepath.Abs(FileCredentialsGCP)

	// drive
	ctx := context.Background()
	// emprego json de service account (ver add_api_token_to_env.sh)
	srv, err = drive.NewService(ctx, option.WithCredentialsFile(FileCredentialsGCP))
	if err != nil {
		slog.Info("Erro drive: " + err.Error())
	}

	// sheets
	ctxSheets := context.Background()
	srvSheets, err = sheets.NewService(ctxSheets, option.WithCredentialsFile(FileCredentialsGCP))
	if err != nil {
		slog.Info("Erro sheets: " + err.Error())
	}

	// calendar
	ctxCal := context.Background()
	cal, err = calendar.NewService(ctxCal, option.WithCredentialsFile(FileCredentialsGCP))
	if err != nil {
		slog.Info("Erro calendar: " + err.Error())
	}
}

// devolve o Id do path pasado, so para directorios
func (d *DirGCP) getDirIdFromName(name string) (string, error) {
	// buscamos polo nome do recurso o
	search, err := srv.Files.List().Q("name='" + name + "' and mimeType='application/vnd.google-apps.folder' and trashed=false").
		Fields("files(id, name, mimeType)").OrderBy("name").Do()
	if err != nil {
		return "", err
	}

	// so debería devolver un resultado
	if len(search.Files) != 1 {
		return "", errors.New("Erro buscando directorio: len(search.Files) = " + strconv.Itoa(len(search.Files)))
	}

	return search.Files[0].Id, nil
}

// lista subdirectorios do path pasado
func (d *DirGCP) listDirectories(path string) ([]string, error) {
	// buscamos Id de path
	id, err := d.getDirIdFromName(path)
	if err != nil {
		return nil, err
	}

	// listamos
	r, err := srv.Files.List().Q("'" + id + "' in parents and mimeType='application/vnd.google-apps.folder' and trashed=false").
		Fields("files(id, name, mimeType)").OrderBy("name").Do()
	if err != nil {
		return nil, err
	}

	var dirs []string
	if len(r.Files) > 0 {
		for _, i := range r.Files {
			dirs = append(dirs, i.Name)
		}
	}
	sort.Strings(dirs)

	return dirs, nil
}

// lista SO ficheiros do path pasado
func (d *DirGCP) listFiles(path string) ([]string, error) {
	slog.Debug("em gcp.listFiles()")

	// buscamos Id de path
	id, err := d.getDirIdFromName(path)
	if err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("id de %s: %s", path, id))

	// listamos
	r, err := srv.Files.List().Q("'" + id + "' in parents and mimeType != 'application/vnd.google-apps.folder' and trashed=false").
		Fields("files(id, name, mimeType)").OrderBy("name").Do()
	if err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("r: %#v", r))

	var files []string
	if len(r.Files) > 0 {
		for _, i := range r.Files {
			// slog.Debug(fmt.Sprintf("file (%#v) ==> (%s) %s", i, i.Id, i.Name)
			files = append(files, i.Name)
		}
	}
	sort.Strings(files)

	return files, nil
}

// busca ficheiro por "name" (nome do ficheiro, nom fai falha completo) e dentro dum cartafol "pathId"
//
//	devolve drive.File do ficheiro atopado
func (d *DirGCP) searchFileNameandDirectory(name string, pathId string) (*drive.File, error) {
	slog.Debug("entrando em searchFile()")

	// buscamos Id de path
	id, err := d.getDirIdFromName(pathId)
	if err != nil {
		return nil, err
	}
	slog.Debug("recolhido ID de " + pathId + ": " + id)

	// buscamos polo nome do recurso
	query := "name contains '" + name + "' and '" + id + "' in parents and mimeType != 'application/vnd.google-apps.folder' and trashed=false "
	slog.Debug("Query de busqueda em GCP: " + query)

	search, err := srv.Files.List().Q(query).Fields("files(id, name, mimeType)").OrderBy("name").Do()
	if err != nil {
		return nil, err
	}

	// so debería devolver un resultado
	if len(search.Files) != 1 {
		slog.Debug("erro, atopado mais (ou menos) dum ficheiro: total atopados = " + strconv.Itoa(len(search.Files)))
		return nil, errors.New("Erro buscando directorio: len(search.Files) = " + strconv.Itoa(len(search.Files)))
	}
	slog.Debug("atopo 1 so ficheiro: (" + search.Files[0].Id + ") " + search.Files[0].Name)

	return search.Files[0], nil
}

// devolve o Id do ficheiro que comeza polo nome pasado (por exemplo para documentos con fecha, por "YYYY-MM-DD") e no path indicado
func (d *DirGCP) getFile(name string, pathId string) ([]byte, error) {
	slog.Debug("entrando em getFile()")

	file, err := d.searchFileNameandDirectory(name, pathId)
	if err != nil {
		return nil, err
	}

	// baixamos, é umha http.Response, asi que lemos para convertir em []byte
	httpRes, err := srv.Files.Get(file.Id).Download()
	if err != nil {
		return nil, err
	}

	defer httpRes.Body.Close()

	fileLocal, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	slog.Debug("gardado ficheiro em []byte")

	return fileLocal, nil
}

// Devolve o path absoluto a um ficheiro temporal descargado, previamente EXPORTADO e descargado em DIRECTORIO_TEMPORAL
//
//	Recibe os mesmo parametros que getFile()+mimetype:
//		name (nome do ficheiro, nom tem porque ser todo o nome enteiro),  pathId (ID do cartafol que contém o ficheiro) e mimetype
func (d *DirGCP) exportAndDownloadFile(name string, pathId string, mimetype string) (string, error) {
	slog.Debug("entrando em exportAndDownloadFile()")

	file, err := d.searchFileNameandDirectory(name, pathId)
	if err != nil {
		return "", err
	}

	// se nom exportamos, descargamos o ficheiro orixinal, polo que creamos umha segunda variable
	extension := ""
	ext, _ := mime.ExtensionsByType(mimetype)
	if len(ext) > 0 {
		extension = ext[0]
	}
	fileName := strings.TrimSuffix(file.Name, filepath.Ext(file.Name)) + extension
	pathAbsoluto := os.TempDir() + string(os.PathSeparator) + fileName
	slog.Debug("pathAbsoluto " + pathAbsoluto)

	// buscamos se o ficheiro esta ja descargado, ou exportado ou o orixinal
	if fi, err := os.Stat(pathAbsoluto); err == nil {
		// comprobamos tamanho porque se tentamos descargalo antes e dá erro, queda creado um ficheiro de 0 bytes (ver providers.GetFileOrExportToPDFLocal() )
		if fi.Size() > 0 {
			slog.Debug("existe ficheiro " + pathAbsoluto + ", nom fai falha descargalo")
			slog.Debug(fmt.Sprintf("fi.size ficheiro: %d", fi.Size()))

			return pathAbsoluto, nil
		}
	}

	slog.Debug("A descargar ficheiro: (" + file.Id + " - " + file.Name + ") " + fileName)

	// exportar
	httpRes, err := srv.Files.Export(file.Id, mimetype).Download()
	if err != nil {
		slog.Debug("Erro exportando e descargando ficheiro: (" + err.Error())
		return "", err
	}

	out, err := os.Create(pathAbsoluto)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, httpRes.Body)
	if err != nil {
		return "", err
	}
	defer httpRes.Body.Close()

	slog.Debug("Descarga OK: (" + file.Id + " - " + file.Name + ") " + fileName)

	return pathAbsoluto, nil
}

// Devolve o path absoluto a um ficheiro temporal descargado, previamente descargado em DIRECTORIO_TEMPORAL
//
//	Recibe os mesmo parametros que getFile():
//		name (nome do ficheiro, nom tem porque ser todo o nome enteiro),  pathId (ID do cartafol que contém o ficheiro)
func (d *DirGCP) downloadFile(name string, pathId string) (string, error) {
	slog.Debug("entrando em downloadFile()")

	file, err := d.searchFileNameandDirectory(name, pathId)
	if err != nil {
		return "", err
	}

	// se nom exportamos, descargamos o ficheiro orixinal, polo que creamos umha segunda variable
	pathAbsoluto := os.TempDir() + string(os.PathSeparator) + file.Name
	slog.Debug("pathAbsoluto " + pathAbsoluto)

	// buscamos se o ficheiro esta ja descargado, ou exportado ou o orixinal
	if fi, err := os.Stat(pathAbsoluto); err == nil {
		// comprobamos tamanho porque se tentamos descargalo e tem 0 bytes, peta
		if fi.Size() > 0 {
			slog.Debug("existe ficheiro " + pathAbsoluto + ", nom fai falha descargalo")
			slog.Debug(fmt.Sprintf("fi.size ficheiro: %d", fi.Size()))

			return pathAbsoluto, nil
		}
	}
	slog.Debug("A descargar ficheiro: " + file.Id + " - " + file.Name)

	// Tratamos de baixar. Se peta, tratamos de exportar antes
	//  todo isto vem porque nom existe propiedade que indique se o arquivo é exportable ou nom de xeito doado
	httpRes, errGet := srv.Files.Get(file.Id).Download()
	if errGet != nil {
		if err != nil {
			slog.Debug("Erro exportando e descargando ficheiro: " + err.Error())
			return "", err
		}
	}

	out, err := os.Create(pathAbsoluto)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, httpRes.Body)
	if err != nil {
		return "", err
	}
	defer httpRes.Body.Close()

	slog.Debug("Descarga OK: " + file.Id + " - " + file.Name)

	return pathAbsoluto, nil
}

// Devolve []byte dumha spreadsheet
//
//	Recibe os mesmo parametros que getFile():
//		name (nome do ficheiro, nom tem porque ser todo o nome enteiro),  pathId (ID do cartafol que contém o ficheiro)
func (d *DirGCP) getSpreadsheetTab(name string, pathId string, tabAndRange string) ([][]interface{}, error) {
	slog.Debug("entrando em getSpreadsheet()")

	file, err := d.searchFileNameandDirectory(name, pathId)
	if err != nil {
		return nil, err
	}
	slog.Debug("file.Id: (" + file.Id + ")")

	// empregamos service de spreadsheet
	resp, err := srvSheets.Spreadsheets.Values.Get(file.Id, tabAndRange).Do()
	if err != nil {
		return nil, err
	}

	var values [][]interface{}

	if len(resp.Values) == 0 {
		// fmt.Println("No data found.")
		return values, nil
	}

	// para crear link despois de consumido
	IdFile = file.Id

	return resp.Values, nil
}
