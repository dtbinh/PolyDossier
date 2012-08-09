// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"flag"
	"log"
	"net/http"
	//"studash/adapters" // uncomment for test 
	"os"
	"studash/tools"
)

var (
	debug = flag.Bool("d", false, "Partir le program en mode debug")
	dir   = flag.String("client", WorkDirError, "Dossier client [obligatoire]")
)

// Le nom à contacter pour le dossier étudiant.
const PolyHostName = "https://www4.polymtl.ca"
const WorkDirError = ""

// Function init nous permettant d'initialiser le débogging.
func init() {
	flag.Parse()
	if *dir == WorkDirError {
		flag.PrintDefaults()
		os.Exit(1)
	}
	tools.SetUpLogging(*debug)
}

// Point d'entré de notre program studash.
func main() {
	// test the ADAPTORS
	//ab := []adapters.AdapterBuilder{adapters.InformationBuilder{}, adapters.ReportBuilder{}, adapters.ScheduleBuilder{}, adapters.InternshipBuilder{}, adapters.FrequentationBuilder{}, adapters.FinanceBuilder{}}
	//for _, adapter := range ab {
		//ab[0].GetParser().GetValue("", "", tools.HTMLParameter{})
	//}
	// end of test

	log.Println("[INFO] : Démarrage Serveur")
	http.HandleFunc("/", onHandleRequest)
	err := http.ListenAndServe("127.0.0.1:http", nil)

    if err != nil { 
      log.Println(err)
      panic(err)
    }
	tools.ExitSuccess()
}
