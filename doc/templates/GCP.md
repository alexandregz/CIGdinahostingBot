# Habilitar *provider* GCP

Para activar o proveedor de Google Cloud Platform como *backend*, podemos ir debugueando co package *temas*, executando os tests no package *temas*.

## Test *temas*


1. Se probamos por primeira vez, sen ter dado acceso ao proxecto e/ou sen habilitar a API de Google Drive, executará o seguinte test con erro:  

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/temas(master)$ go test -test.v
=== RUN   TestMenuTemasGCP
2023/11/04 23:30:39 Erro: googleapi: Error 403: Google Drive API has not been used in project 19xxxxxxx before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/drive.googleapis.com/overview?project=19xxxxxx then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
Details:
[
  {
    "@type": "type.googleapis.com/google.rpc.Help",
    "links": [
      {
        "description": "Google developers console API activation",
        "url": "https://console.developers.google.com/apis/api/drive.googleapis.com/overview?project=19xxxxxx"
      }
    ]
  },
  {
    "@type": "type.googleapis.com/google.rpc.ErrorInfo",
    "domain": "googleapis.com",
    "metadata": {
      "consumer": "projects/19xxxxx",
      "service": "drive.googleapis.com"
    },
    "reason": "SERVICE_DISABLED"
  }
]
, accessNotConfigured
exit status 1
FAIL	github.com/alexandregz/CIGdinahostingBot/src/temas	1.487s
```

```go
alex@vosjod:~/Development/CIGdinahostingBot/src/temas(master)$ go test -test.v --run ^TestMenuTemasGCPxlsx$
=== RUN   TestMenuTemasGCPxlsx
2023/11/05 00:25:03 Erro: googleapi: Error 403: Google Sheets API has not been used in project 195871835206 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/sheets.googleapis.com/overview?project=195871835206 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
Details:
[
  {
    "@type": "type.googleapis.com/google.rpc.Help",
    "links": [
      {
        "description": "Google developers console API activation",
        "url": "https://console.developers.google.com/apis/api/sheets.googleapis.com/overview?project=195871835206"
      }
    ]
  },
  {
    "@type": "type.googleapis.com/google.rpc.ErrorInfo",
    "domain": "googleapis.com",
    "metadata": {
      "consumer": "projects/195871835206",
      "service": "sheets.googleapis.com"
    },
    "reason": "SERVICE_DISABLED"
  }
]
, accessNotConfigured
exit status 1
FAIL	github.com/alexandregz/CIGdinahostingBot/src/temas	2.119s
```


NOTA: Se non temos activada a *Google Cloud Console* na nosa conta, primeiro debemos activala, introducindo os datos persoais que nos indica e a tarxeta de débito/crédito que empregan para comprobar que se poden realizar cobros (só nos cobrarán por consumo, dándonos 300$ para probas gratis).




2. Habilitamos vía web as API de *Google Drive* e de *Google Sheets* (ver a mensaxe *Enable it by visiting https://console.developers.google.com/apis/api/drive.googleapis.com/overview?project=19xxxxxx*):

![Habilitar Google Drive API](../img/conta_servizo_GCP_test_real/001_GCP_API_Drive_habilitar.png)

![Habilitar Google Sheets API](../img/conta_servizo_GCP_test_real/001_GCP_API_Sheets_habilitar.png)



3. Habilitada a API, o problema é que non atopa o ficheiro:

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/temas(master)$ go test -test.v
=== RUN   TestMenuTemasGCP
2023/11/04 23:37:20 Erro: Erro buscando directorio: len(search.Files) = 0
exit status 1
FAIL	github.com/alexandregz/CIGdinahostingBot/src/temas	2.214s
```

4. Engadimos ficheiros ao noso espazo de Google Drive.

Se comprobamos os tests (*temas/temas_test.go*), empregan as seguintes funcións do package *providers*:

```go
	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFileGCPtxt("Temas en curso LISTADO", "Temas en curso")

	// GetFileGCP(file, directory), devolve byte de ler o ficheiro exportado e descargado
	b, err := providers.GetFileGCPxlsx("Temas en curso LISTADO", "Temas en curso", "Temas en curso")
```

Isto quere dicir que os ficheiros que empregan **están nun directorio *Temas en curso*** e dentro existe unha folla de cálculo de nome **Temas en curso LISTADO**, cunha pestaña de nome **Temas en curso**


5. **IMPORTANTE**: Temos que compartir o directorio coa conta de servizo creada para que teña, cando menos, permisos de lectura, se non segue dando erro que non atopa os ficheiros/directorio.

![Compartir directorio coa conta de servizo](../img/conta_servizo_GCP_test_real/005_compartir_directorio.png)


6. Confirmamos acceso ao ficheiro de *Temas en curso LISTADO* co test:

```go
alex@vosjod:~/Development/CIGdinahostingBot/src/temas(master)$ go test -test.v --run ^TestMenuTemasGCPxlsx$
=== RUN   TestMenuTemasGCPxlsx
b: [[][]interface {}{[]interface {}{"Data comezo", "Data fin", "Tema", "Descripcion"}, []interface {}{"", "", "revisar xx xx 2022", "Dixo Ester de Igualdade da CIG que xxx."}, []interface {}{"", "2023-04-21", "", "Postura a tomar diante dos contratos xxx: agardamos a nova contratación em soporte, a ver xxx."}, []interface {}{"", "", "reunións previas ao comité", "realizar reunións previas entre a sección sindical previas á reunión do comité de empresa"}, []interface {}{"", "2023-05-01", "cláusula proibidos xxx xxxx", "para o CSS, realizar cláusula para proibir xxxxxx, que se engada ao Protocolo de Desconexión dixital"}}]--- PASS: TestMenuTemasGCPxlsx (1.62s)
PASS
ok  	github.com/alexandregz/CIGdinahostingBot/src/temas	2.599s
```