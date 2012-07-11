// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"net/http"
	"studash/tools"
	"log"
	"flag"
)

var ( trololo
	debug = flag.Bool("d", false, "Partir le program en mode debug")
)

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

	tools.ExitSuccess();
}