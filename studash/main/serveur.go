// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"studash/errors"
	//"studash/pages"
	"time"
)

func onHandleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	doHandleRequest(w, r)
	log.Println(fmt.Sprintf("[%s] - %s", time.Since(start), r.URL.Path))
}

func doHandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := strings.Split(r.URL.Path, "/")

	// Page root 
	// Note : Étant donné que le premier charactere est un slash, il donne vide en 0.
	if ctx[1] == "" {
		w.Write(DefaultPage())
		return
	}

	// Reste du site
	data, err := do(r, ctx)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	w.Write(data)
}

func do(r *http.Request, ctx []string) ([]byte, error) {
	var data []byte = nil
	switch ctx[1] {
	case "auth":
		data = []byte("Authentification")
	case "u":
		data = []byte("Utilisateur")
	case "c":
		data = []byte("Comites")
	case "p":
		data = []byte("Poly")
	}
	if data != nil {
		return data, nil
	}
	return nil, &errors.RequestError{Action: ctx[1], Method: r.Method, Problem: "Action illegale."}
}

func DefaultPage() []byte {
	file, err := os.Open("../src/studash/client/index.html")
	defer file.Close()
	if err != nil {
		log.Print("[ERROR] : Def Pag Open", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("[ERROR] : Def Pag Read", err)
	}

	return data
}
