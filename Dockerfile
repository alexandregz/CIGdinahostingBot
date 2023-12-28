# syntax=docker/dockerfile:1

# docker build -t cigdhbot-docker:multistage -f Dockerfile .

# notas para docker run
# 1) empregamos a rede do host
# 2) montamos localmente os ficheiros de config de usuarios e da chave privada do cert (por seguridade evidentemente NON quedan na imaxe)
# 3) forzamos o endpoint para poder pasar, despois da imaxe, os parámetros via argumentos de cli (ao empregar distroless non temos environment)
#
# docker run --network host -v "$(pwd)"/config.yml:/tmp/config.yml -v "$(pwd)"/service-account.json:/tmp/service-account.json --entrypoint /cigdhbot-docker cigdhbot-docker:multistage '-token' 'x:xx' '-canle' 'xx' '-chatID' '-xx' '-configFileUsers' '/tmp/config.yml' '-gcpCredentialsFile' '/tmp/service-account.json' 
#

FROM golang:1.21 AS build-stage

WORKDIR /app

# directorio para deixar o binario e os ficheiros de config
RUN mkdir bin

ADD src .
WORKDIR /app/src
RUN go mod download
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/bin/cigdhbot-docker main.go



# # # Run the tests in the container
# FROM build-stage AS run-test-stage
# WORKDIR /app
# RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app/bin

COPY --from=build-stage /app/bin/cigdhbot-docker /cigdhbot-docker

ENTRYPOINT ["/cigdhbot-docker"]
