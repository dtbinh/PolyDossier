// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"studash/errors"
	"studash/pages"
	"studash/tools"
)

var (
	debug = flag.Bool("d", false, "Partir le program en mode debug")
)

const kRoot = "/"

// Le nom à contacter pour le dossier étudiant.
const PolyHostName = "https://www4.polymtl.ca"

// Function init nous permettant d'initialiser le débogging.
func init() {
	flag.Parse()
	tools.SetUpLogging(*debug)
}

// Point d'entré de notre program studash.
func main() {
	log.Println("[INFO] : Démarrage Serveur")
	http.HandleFunc("/", onHandleRequest)
	http.ListenAndServe(":http", nil)

	tools.ExitSuccess()
}

func doAction(method, path string) ([]byte, error) {
	switch method {
	case "GET":
		return doGet(path)
	case "POST":
		return doPost(path)
	case "PUT":
		return doPut(path)
	case "DELETE":
		return doDelete(path)
	}
	return []byte{}, errors.ErrUnimplemented
}

func doGet(path string) ([]byte, error) {
	c := pages.Credentials{"temp", "temp", "temp"}

	if path[0:3] == "/u/" {
		if strings.Contains(path[4:], "/") {
			endidx := strings.LastIndex(path[4:], "/")
			return pages.FetchInfos(c, path[5+endidx:]), nil
		} else {
			return pages.ListFunctions(c), nil
		}

	}

	return []byte{}, errors.ErrUnimplemented
}
func doPost(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}
func doPut(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}
func doDelete(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}
