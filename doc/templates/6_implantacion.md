# FASE DE IMPLANTACIÓN

## Manual técnico:

### Información relativa á instalación e programación: 

#### Requerimentos

* Servidor **nix*: GNU Linux/Unix
* Backend:
   * Acceso á [consola de GCP](https://console.cloud.google.com/) para poder dar de alta unha *conta de servizo* cunha *chave privada.*


##### Para desenvolvemento

* Versión de *go* desexable 1.17 (a nova API de Go para Google Cloud require esa versión mínima, polo de agora non está portada pero se queremos empregar esa nova API, como aconsella Google, sería o axeitado)
* Ficheiro YAML `config.yml` para definir usuarios e niveis destes.
* Ficheiro `service_account.json` de autenticación en Google Cloud (para empregar GCP como *backend*)

Exemplo de `config.yml`:

```yaml
users:
- 1234567   #Alex
- 1234568   # test
admins:
- 1234567
superadmins:
- 1234567
```

Os IDs de usuarios hai que extraelos manualmente. Pódese empregar o bot de Telegram [@userinfobot](https://t.me/userinfobot)


Exemplo de `service_account.json`:

```json
{
  "type": "service_account",
  "project_id": "tempxxx-xxxxx-xxxxxx",
  "private_key_id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDXBGF5+7YLgAFC\nnhSSV4T
  
  [...]
  
  /nNN3Sw\now6f39IZ+p9LvFoaFR92jGs=\n-----END PRIVATE KEY-----\n",
  "client_email": "conta-servizo-cigdhbot@tempxx-xxx-xxxxxx.iam.gserviceaccount.com",
  "client_id": "xxxxxxxxxxxxxxxxx",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/conta-servizo-cigdhbot%40tempxx-xxxx-xxxxxxxxxx.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
```

##### Crear conta de servizo de GCP con chave privada

Titorial sobre creación de conta de servizo para empregar como identificador chaves privadas desta.

[Crear conta de servizo en GCP](7_crear_service_account_GCP.md)


##### Inicializar módulos

Para inicializar os módulos de *go* e poder ir realizando un seguemento das dependencias do código, emprego **go mod**

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src$ go mod init github.com/alexandregz/CIGdinahostingBot/src
go: creating new go.mod: module github.com/alexandregz/CIGdinahostingBot/src
go: to add module requirements and sums:
	go mod tidy
alex@vosjod:~/Development/CIGdinahostingBot/src$ go mod tidy
go: finding module for package golang.org/x/exp/slices
go: finding module for package google.golang.org/api/option
go: finding module for package gopkg.in/yaml.v2
go: finding module for package google.golang.org/api/sheets/v4
go: finding module for package modernc.org/sqlite
go: finding module for package github.com/go-telegram-bot-api/telegram-bot-api/v5
go: finding module for package google.golang.org/api/drive/v3
go: finding module for package gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud
go: found github.com/go-telegram-bot-api/telegram-bot-api/v5 in github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
go: found golang.org/x/exp/slices in golang.org/x/exp v0.0.0-20231226003508-02704c960a9b
go: found gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud in gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud v1.2.3
go: found google.golang.org/api/drive/v3 in google.golang.org/api v0.154.0
go: found google.golang.org/api/option in google.golang.org/api v0.154.0
go: found google.golang.org/api/sheets/v4 in google.golang.org/api v0.154.0
go: found modernc.org/sqlite in modernc.org/sqlite v1.28.0
go: found gopkg.in/yaml.v2 in gopkg.in/yaml.v2 v2.4.0
```




### Implantación

Para poñer en marcha o proxecto:

1. Crear conta de servizo
2. Crear bot en Telegram recollendo token
3. Empregar script `add_api_token_to_env.sh`, engadindo os valores:
	- `BOT_TG_CIGDH_APITOKEN`: Token de Telegram do bot
	- `BOT_TG_CIGDH_CANLE_PROBAS_DEFAULT`: Canle de probas (desenvolvemento)
	- `BOT_TG_CIGDH_CANLE_ID_DEFAULT`: Canle por defecto (producción)
	- `CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP`: chave privada de conta de servizo de GCP. Full path á súa ruta.
	- `FILE_CONFIG_USERS`: path absoluto ao ficheiro yaml de configuración de usuarios


#### Documentación de GCP

Se quere empregarse GCP como backend (desexado), cómpre revisar a documentación de *Google Sheets API* e *Google Drive API*: 

- [https://developers.google.com/sheets/api/quickstart/go](https://developers.google.com/sheets/api/quickstart/go)
- [https://developers.google.com/drive/api/quickstart/go](https://developers.google.com/drive/api/quickstart/go)


Golang APIs Google Sheets e Google Drive:

- [https://pkg.go.dev/google.golang.org/api@v0.111.0/drive/v3](https://pkg.go.dev/google.golang.org/api@v0.111.0/drive/v3)

- [https://pkg.go.dev/google.golang.org/api@v0.111.0/sheets/v4](https://pkg.go.dev/google.golang.org/api@v0.111.0/sheets/v4)

- [https://pkg.go.dev/google.golang.org/api/option](https://pkg.go.dev/google.golang.org/api/option)



### Administración do sistema

Emprégase *systemd* para *daemonizar* o binario xerado. Empregamos un ficheiro  para crear a *unit* do servizo (*cigdhbot.service*), que apunta a un directorio onde residen:

- O ficheiro da chave cifrada de Google Cloud: *service-account.json*
- O ficheiro de config de usuarios autorizados: *config.yml*
- Un ficheiro contendo as variables de entorno a cargar na *unit*, definido na key *EnvironmentFile* de *Service* (ver *cigdhbot.service*)


*cigdhbot.service*

```bash
[Unit]
Description=CLI CIGdhBot
After=network.target

[Service]
User=pi
Type=simple

WorkingDirectory=/home/pi/FP_SAN_CLEMENTE_proxecto_DAM_distancia
EnvironmentFile=/home/pi/FP_SAN_CLEMENTE_proxecto_DAM_distancia/add_api_token_to_env_systemd.sh
ExecStart=/bin/sh -c "/home/pi/FP_SAN_CLEMENTE_proxecto_DAM_distancia/cigdhbot_raspberry -debug Debug >> /var/log/cigdhbot/cigdhbot.log 2>&1"
StandardError=journal

Restart=on-failure
RestartSec=5
Nice=19

#MemoryMax=30M
#CPUQuota=50%

[Install]
WantedBy=multi-user.target
```

*EnvironmentFile*

```bash
# https://t.me/CIGdhBot
BOT_TG_CIGDH_APITOKEN=xxxx:xxxxxxxx

BOT_TG_CIGDH_CANLE_PROBAS_DEFAULT=@xxxx

BOT_TG_CIGDH_CANLE_ID_DEFAULT=-xxxxxx

# se empregamos de provider Google Cloud. Olho! fullpath
CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP=/home/pi/FP_SAN_CLEMENTE_proxecto_DAM_distancia/service-account.json

# yaml config de usuarios, tamem fullpath
FILE_CONFIG_USERS=/home/pi/FP_SAN_CLEMENTE_proxecto_DAM_distancia/config.yml
```



### Backend

Emprégase Google Drive como backend dos datos.

Para abstraer o consumo de datos e poder reemprazar de xeito doado o backend por outro, empregouse o patrón Template Design ([https://refactoring.guru/design-patterns/template-method](https://refactoring.guru/design-patterns/template-method)) e asi poder reemprazar de xeito doado o backend se fixera falla.

A implementación está no paquete `providers`, pois son os métodos que proveen ao resto das funcionalidades que se consumen:

- Ler ficheiros dun directorio (GCP)
- Baixar ficheiros (GCP)
- Exportar ficheiros e baixalos (GCP)
- Ler e escribir en follas de cálculo (GCP)

En `gcp.go` creanse os métodos particulares de Google Cloud Plataform empregados. Para poder acceder aos directorios e ficheiros que se consumen neste `provider`, compartense vía `conta de servizo`, habilitando as APIs necesarias para o seu consumo: `Google Drive API` e `Google Sheets API`.
A conta de servizo identíficase pasando como variable de entorno (`CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP`) a ruta ao ficheiro `service-account.json` descargable desde a consola de Google Cloud. Ver exemplo en `EXAMPLE add_api_token_to_env.sh`.

En `localfilesystem.go` existen as implementacións dos métodos para empregar un sistema de ficheiros local, polo de agora sen uso.

En `sqlite.go` impleméntase o consumo dun ficheiro SQLite.

En `nextcloud.go` impleméntase un exemplo de emprego dun Nextcloud como backend, empregando o package [https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud](https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud) como acceso a esta API.

[Documentación de *providers*](providers.md)



## Tests

#### Test completo

Inclúe cobertura de código e comprobación de *race conditions*:

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src(master)$ source ../add_api_token_to_env.sh &&  go test -cover -race ./...
?   	github.com/alexandregz/CIGdinahostingBot/src	[no test files]
?   	github.com/alexandregz/CIGdinahostingBot/src/canle	[no test files]
?   	github.com/alexandregz/CIGdinahostingBot/src/comite	[no test files]
?   	github.com/alexandregz/CIGdinahostingBot/src/handlers	[no test files]
?   	github.com/alexandregz/CIGdinahostingBot/src/menus	[no test files]
?   	github.com/alexandregz/CIGdinahostingBot/src/lexislacion	[no test files]
ok  	github.com/alexandregz/CIGdinahostingBot/src/bot	2.132s	coverage: 88.9% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/comandos	2.340s	coverage: 92.9% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/providers	6.853s	coverage: 31.6% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/senders	2.070s	coverage: 18.4% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/temas	7.001s	coverage: 0.0% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/user	3.221s	coverage: 79.2% of statements
ok  	github.com/alexandregz/CIGdinahostingBot/src/utils	11.953s	coverage: 51.7% of statements
```


#### Providers

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/providers(master)$ go test -test.v
=== RUN   TestGetFileOrExportToPDFLocal
Descargado: [/var/folders/mx/tgfkvtc57_qbzsv0hm_m33z80000gn/T//2021-01-13 ACTA DE LA COMISIÓN DE IGUALDAD.pdf]
--- PASS: TestGetFileOrExportToPDFLocal (1.54s)
=== RUN   TestGetFileOrExportToPDFLocal2
ficheiro1 em local (Exportado previamente): /var/folders/mx/tgfkvtc57_qbzsv0hm_m33z80000gn/T//acta 26-01-23.pdf
ficheiro2 em local (so descarga): /var/folders/mx/tgfkvtc57_qbzsv0hm_m33z80000gn/T//2021-01-13 ACTA DE LA COMISIÓN DE IGUALDAD.pdf
--- PASS: TestGetFileOrExportToPDFLocal2 (3.01s)
=== RUN   TestGetListDirectoriesNextcloud
--- PASS: TestGetListDirectoriesNextcloud (2.59s)
=== RUN   TestGetListFilesNextcloud
--- PASS: TestGetListFilesNextcloud (2.44s)
PASS
ok  	github.com/alexandregz/CIGdinahostingBot/src/providers	10.677s
```

Debug de Nextcloud: 
```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/providers(master)$ source ../../add_api_token_to_env.sh && go test -run ^TestGetListFilesNextcloud$
0 - Nextcloud Manual.pdf
1 - Nextcloud intro.mp4
2 - Nextcloud.png
3 - Readme.md
4 - Reasons to use Nextcloud.pdf
5 - Templates credits.md
PASS
ok  	github.com/alexandregz/CIGdinahostingBot/src/providers	2.880s
```

#### Temas

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/temas(master)$ go test -test.v
=== RUN   TestMenuTemasGCP
str: [tema1
tema2
e
tema4]--- PASS: TestMenuTemasGCP (1.49s)
=== RUN   TestMenuTemasGCPxlsx
b: [[][]interface {}{[]interface {}{"Data comezo", "Data fin", "Tema", "Descripcion"}, []interface {}{"", "", "revisar xx xx 2022", "Dixo Ester de Igualdade da CIG que xxx."}, []interface {}{"", "2023-04-21", "", "Postura a tomar diante dos contratos xxx: agardamos a nova contratación em soporte, a ver xxx."}, []interface {}{"", "", "reunións previas ao comité", "realizar reunións previas entre a sección sindical previas á reunión do comité de empresa"}, []interface {}{"", "2023-05-01", "cláusula proibidos xxx xxxx", "para o CSS, realizar cláusula para proibir xxxxxx, que se engada ao Protocolo de Desconexión dixital"}}]--- PASS: TestMenuTemasGCPxlsx (1.09s)
PASS
ok  	github.com/alexandregz/CIGdinahostingBot/src/temas	3.304s
```

#### Utils

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src/utils(master)$ go test -test.v
=== RUN   TestSentenzasConformidade
Veña.
👌🏻

Entendido.
Entendido.
Entendido.
--- PASS: TestSentenzasConformidade (0.00s)
=== RUN   TestSentenzasDisconformidade
Perdoa.
🧐
Ups.
Perdoa.
Perdoa.
--- PASS: TestSentenzasDisconformidade (0.00s)
--- PASS: TestSentenzasDisconformidade (0.00s)
=== RUN   TestLinksGCP
Baixado xlsx de Temas e exportado a PDF. ContentType: PDF document, version 1.4
--- PASS: TestLinksGCP (3.25s)
=== RUN   TestGetFileContentType
--- PASS: TestGetFileContentType (0.02s)
PASS
ok  	github.com/alexandregz/CIGdinahostingBot/src/utils	2.204s
```


## Librarías Go empregadas

- Go Telegram Bot API: https://go-telegram-bot-api.dev/
- Google APIs Client Library for Go: https://pkg.go.dev/google.golang.org/api
- Cliente de consumo da API de Nextcloud: https://gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud
- Libraría YAML: https://pkg.go.dev/gopkg.in/yaml.v2
- Libraría SQLite para Go (sen CGO): https://pkg.go.dev/modernc.org/sqlite