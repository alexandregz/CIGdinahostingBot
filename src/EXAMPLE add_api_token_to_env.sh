#!/bin/bash

# https://t.me/CIGdhBot
export BOT_TG_CIGDH_APITOKEN='XXX'

export BOT_TG_CIGDH_CANLE_PROBAS_DEFAULT='@XXX'

# id de canle por defecto, podese sobreescribir com parámetro "-chatID '-xxxxxxxxxx'"
#
# nom o dim recolhido bem numha canle ata que falou outro usuario do grupo (com "-debug API")
# recolhido via https://web.telegram.org/ daba um chatID distinto
export BOT_TG_CIGDH_CANLE_ID_DEFAULT='-XXXXXXXXXXX'


# se empregamos de provider Google Cloud. Olho! fullpath
export CREDENTIALS_FILE_ACCOUNT_SERVICE_GCP=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )'/XXX'


# yaml config de usuarios, tamem fullpath
export FILE_CONFIG_USERS=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )'/config.yml'