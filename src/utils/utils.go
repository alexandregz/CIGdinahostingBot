// utilidades varias
package utils

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

// devolve aleatoriamente umha sentencia de conformidade, para "humanizar" as respostas dos menús, mensagens de infor, etc.
func SentenzasConformidade() string {

	sentenzas := []string{
		"De acordo.",
		"Moi ben.",
		"👌🏻 \n",
		"Vale.",
		"Veña.",
		"Perfecto.",
		"Entendido.",
	}

	rand.New(rand.NewSource(time.Now().Unix()))
	return sentenzas[rand.Intn(len(sentenzas))]
}

// devolve aleatoriamente umha sentencia de disconformidade
func SentenzasDisconformidade() string {

	sentenzas := []string{
		"Perdoa.",
		"🧐  ",
		"Que pasou?.",
		"Ups.",
	}

	rand.New(rand.NewSource(time.Now().Unix()))
	return sentenzas[rand.Intn(len(sentenzas))]
}

// fai pad de string a tamanho "s", devolvendo um slice co string padeado
func ChunkSplit(body string, limit int, end string) string {

	// comprobamos que faga falha chunkear
	if len(body) <= limit {
		return body
	}

	var charSlice []rune

	// push characters to slice
	for _, char := range body {
		charSlice = append(charSlice, char)
	}

	var result string = ""

	for len(charSlice) >= 1 {
		// convert slice/array back to string
		// but insert end at specified limit

		result = result + string(charSlice[:limit]) + end

		// discard the elements that were copied over to result
		charSlice = charSlice[limit:]

		// change the limit
		// to cater for the last few words in
		// charSlice
		if len(charSlice) < limit {
			limit = len(charSlice)
		}

	}

	return result
}

// devolve os links de documentos de google drive, para empregalos com Sprintf()
func LinksGCP(tipo string, id string) string {

	links := map[string]string{

		"spreadsheet": "https://docs.google.com/spreadsheets/d/%s/",
		"doc":         "https://docs.google.com/document/d/%s/",
		"form":        "https://docs.google.com/forms/d/%s/",
	}

	val, ok := links[tipo]
	if ok {
		return fmt.Sprintf(val, id)
	}

	return ""
}

// devolve o filetype dun ficheiro (inclue o path, relativo ou absoluto a este directorio)
func GetFileContentType(f string) (string, error) {
	out, err := exec.Command("/usr/bin/file", f).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	contentType := strings.Replace(string(out[:]), f+": ", "", 1)
	contentType = strings.TrimSuffix(contentType, "\n")

	return contentType, nil
}
