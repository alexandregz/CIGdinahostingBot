# CIGdhBot

![CIGdhBot logo](doc/img/logo_pequeno.png)

## Descripción


Bot de [Telegram](https://telegram.org/faq/es?setln=es#p-que-es-telegram-que-puedo-hacer-aqui) en [Go](https://go.dev/) para xestionar a información e o funcionamento diario dun Comité de empresa, neste caso da compostelana [dinahosting](https://dinahosting.com/).

A través dunha serie de menús permite a consulta rápida online de información de consulta necesaria para o Comité de empresa/sección sindical, como poden ser *temas en curso*, *histórico de ficheiros*, *consulta de actas* de órganos como o Comité de Seguridade e Saúde (CSS) ou o Comité negociador dun Plan de Igualdade, etc.

## Instalación / Posta en marcha

A aplicación consta dun código en Golang correspondente ao bot e de ficheiros aloxados en Google Drive que se empregarán como *backend* e dos que se extrae a información. Así conseguimos que a información este centralizada nun sitio accesible de múltiples xeitos para un conxunto de persoas cun custo nulo.


Para lanzar a aplicación necesitaremos definir certas variables:

- O *token* que asigna Telegram ao Bot (ver [https://core.telegram.org/bots/tutorial](https://core.telegram.org/bots/tutorial))
- Como empregamos forzosamente Google Cloud como *backend*, necesitamos empregar o exportado de credenciais de GCP (Google Cloud Platform) (ver [Crear conta de servizo en GCP](doc/templates/7_crear_service_account_GCP.md))
- Hai que definir un ficheiro de configuración dos usuarios autorizados a empregar a ferramenta, en formato YAML (ver [Implantación](doc/templates/6_implantacion.md))


Por seguridade e practicidade, seguiremos a metodoloxía de [*12factors*](https://12factor.net/es/) de setear a configuración en variables de entorno: [https://12factor.net/es/config](https://12factor.net/es/config), aínda que finalmente tamén se engadiu a posibilidade de pasar argumentos via CLI:

```bash
alex@vosjod:~/Development/CIGdinahostingBot/src$ source ../add_api_token_to_env.sh  && go run main.go --help
Usage of /var/folders/mx/tgfkvtc57_qbzsv0hm_m33z80000gn/T/go-build4044499318/b001/exe/main:
  -canle string
    	Canle a conectar, por exemplo 'probas_cigdhbot' (default "@asminhasprobas")
  -chatID int
    	ChatID de canle a conectar, por exemplo '-1234567890' (default -xxxxxxxx)
  -configFileUsers string
    	Path ao ficheiro de config de usuarios
  -debug string
    	'Debug', 'Info', 'Warning', 'Error' ou 'API' (debug de tgbotapi)
  -gcpCredentialsFile string
    	Path do ficheiro de clave secreta de GCP
  -token string
    	Token identificativo do bot
```

#### Configuración script inicio


Polo tanto podemos crear un script de *bash* que cargue esas variables no noso entorno (ao empregar Linux/Mac para despliegue) e sexa chamado previamente a lanzar o código que temos da copia do repositorio localmente:


```bash
alex@vosjod:~/Development/CIGdinahostingBot/src$ source ../add_api_token_nv.sh && go run main.go -debug API
``` 


Exemplo de script en bash:

```bash
#!/bin/bash

# https://t.me/CIGdhBot
export BOT_TG_CIGDH_APITOKEN='XXX'

export BOT_TG_CIGDH_CANLE_PROBAS_DEFAULT='@XXX'

# id de canle por defecto, podese sobreescribir com parámetro "-chatID '-xxxxxxxxxx'"
# recolhido via https://web.telegram.org/ daba um chatID distinto
export BOT_TG_CIGDH_CANLE_ID_DEFAULT='-XXXXXXXXXXX'


# se empregamos de provider Google Cloud. Olho! fullpath
export CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )'/XXX'


# yaml config de usuarios, tamem fullpath
export FILE_CONFIG_USERS=$( cd -- "$( dirname -- "${BASH_SOURCE[    0]}" )" &> /dev/null && pwd )'/config.yml'
```

#### Configuración usuarios

Por seguridade, para poder empregar o bot imos definir unha configuración en [YAML](https://es.wikipedia.org/wiki/YAML) onde gardaremos os IDs dos usuarios autorizados para empregar o bot, por poder conter este información sensible da que se debe gardar o debido sixilo. Empregamos YAML pois permite comentarios -identificando univocamente aos/ás usuarios/as- e é lexible facilmente.

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


## Uso

Cada menú pode conter submenús que se van chamando segundo as necesidades cos seus comandos correspondentes, por exemplo:

1. comezamos pedindo `/axuda`
2. listamos os submenús do menú *Igualdade* co comando `/igualdade`
3. listamos as actas de Igualdade co comando `/igualdade actas`
4. descargamos a acta do día 12 de setembro de 2022 co comando `/igualdade acta 2021-07`


Para contactar o meu enderezo de correo electrónico é: [aemenor@gmail.com](aemenor@gmail.com)




## Licencia

[GPLv3](LICENSE)

## Índice

1. [Implantación](doc/templates/6_implantacion.md)
2. [Crear conta de servizo en GCP](doc/templates/7_crear_service_account_GCP.md)
3. [Package *providers*](doc/templates/providers.md)
4. [Configurar Google Cloud Platform](doc/templates/GCP.md)


## Guía de contribución

As contribucións son benvidas, ben como código cun *Pull Request* -empregando sempre unha *rama* separada-, ben como idea a implementar.

Para contactar co autor, pódese empregar o seguinte email: [aemenor@gmail.com](mailto:aemenor@gmail.com)

## Links

- [Telegram APIs](https://core.telegram.org/)
- [Telegram Bots](https://core.telegram.org/bots)
- [Go Telegram Bot API](https://go-telegram-bot-api.dev/)
- [Contas de servizo en Google Cloud (IAM)](https://console.cloud.google.com/iam-admin/serviceaccounts)

- [Facebook da sección TIC da CIG (CIG-TIC)](https://www.facebook.com/cigtic)
- [Twitter/X da sección TIC da CIG (CIG-TIC)](https://twitter.com/galizacig_tic)




