// Fonctions d'utilité courant tel que Panic et Exit.
package tools

import (
	"fmt"
	"os"
	"time"
	"log"
)

const (
	LogFileName = "log"  // Le prefix du fichier de log que nous produisons.
)

var (
	logfile *os.File = nil  // Le handle du fichier de log ouvert.
)

// Permet le détournement du logging dans un fichier séparé pour consultation ultérieur.
func SetUpLogging(isDebug bool) {
    if isDebug == true { return }
	logfile, err := os.OpenFile(
		fmt.Sprintf("%s_%d%s", LogFileName, time.Now().Unix(), ".log"),
		os.O_WRONLY | os.O_CREATE,
		0666)
		
	if err != nil {
	  log.Fatal("tools/base : [FATAL] Nous n'avons pu créé de fichier de log.\n")
	}
	log.SetOutput(logfile)
}

// Permet de quitter en tenant compte du fichier de log.
func ExitSuccess() {
	if logfile != nil { logfile.Close() }
	os.Exit(0)
}