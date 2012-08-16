// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"studash/adapters"
	"studash/errors"
	"studash/student"
	"time"
)

// On time le temps de traitement des requêtes.
func onHandleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	doHandleRequest(w, r)
	log.Println(fmt.Sprintf("[%s] - %s", time.Since(start), r.URL.Path))
}

func doHandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := strings.Split(r.URL.Path, "/")

	// Si l'utilisateur demande une page sans processing 
	if doStaticPages(ctx[1], w, r) {
		return
	}

	// Reste du site
	data, err := do(r, ctx)
	if err != nil {
		log.Println("[ERROR]", err)
	}

	w.Write(data)
}

func doStaticPages(action string, w http.ResponseWriter, r *http.Request) bool {
	var filename string
	switch action {
	case "":
		filename = "index.html"
		break
	case "_":
		filename = "_"
		w.Header().Set("Content-Type", "application/javascript")
		break
	case "i.png":
		w.Header().Set("Content-Type", "image/png")
		filename = "i.png"
		break
	case "b.gif":
		w.Header().Set("Content-Type", "image/gif")
		filename = "b.gif"
		break
	default:
		return false
		break
	}
	w.Write(GetFileData(*dir + filename))
	return true
}

func do(r *http.Request, ctx []string) ([]byte, error) {
	var data []byte = nil
	switch ctx[1] {
	case "auth":
		data = Authenticate(r)
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

func PullCredentials(r *http.Request) {
}

func GetFileData(filename string) []byte {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Print("[ERROR] : GetFileData Open", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("[ERROR] : GetFileData Read", err)
	}

	return data
}

func Authenticate(r *http.Request) []byte {
	resp, err := http.PostForm(PolyHostName+"/servlet/ValidationServlet", student.BuildCredentials(r).UrlForm())
	if err != nil {
		log.Print("[ERROR] : Sending Auth to poly", err)
	}
	defer resp.Body.Close()

	loginBuilder := adapters.LoginBuilder{}
	matricule, token := loginBuilder.GetMatriculeToken(resp.Body)

	if token != "" {
		log.Print("[INFO] : Matricule :", matricule, ", Token :", token)
		json := fmt.Sprintf(`{"AuthResponse" : true, "Matricule" : %s, "Token" : %s}`, matricule, token)
		return []byte(json)
	}
	log.Print("[ERROR] : Invalid Credentials or 'Dossier Etudiant' is down")
	return []byte(`{"AuthResponse" : false}`)
}

func IsReaderAtLeast(r io.Reader, size int) bool {
	tmp := make([]byte, size)

	_, err := io.ReadAtLeast(r, tmp, size)
	if err == io.ErrUnexpectedEOF {
		return false
	}
	return true
}
