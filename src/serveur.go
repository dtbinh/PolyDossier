// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"studash/errors"
	"studash/adapters"
	//"studash/student"
	"io"
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

	// Le JS compiler par Closure
	if ctx[1] == "_" {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(DefaultScript())
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

func DefaultPage() []byte {
	file, err := os.Open(*dir + "index.html")
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

func DefaultScript() []byte {
	file, err := os.Open(*dir + "_")
	defer file.Close()
	if err != nil {
		log.Print("[ERROR] : Def script Open", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("[ERROR] : Def script Read", err)
	}

	return data
}

func Authenticate(r *http.Request) []byte {
	postUrl := PolyHostName + "/servlet/ValidationServlet"

	defer r.Body.Close()
	jsonForm, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("[ERROR] : Json Form ReadAll", err)
	}

	var formValues map[string]string
	err = json.Unmarshal(jsonForm, &formValues)
	if err != nil {
		log.Print("[ERROR] : Json Form UnMarshall", err)
	}

	formz := url.Values{}
	for key, value := range formValues {
		switch key {
		case "username":
			formz.Add("code", value)
		case "password":
			formz.Add("nip", value)
		case "dateOfBirth":
			formz.Add("naissance", value)
		}
	}

	resp, err := http.PostForm(postUrl, formz)

	if err != nil {
		log.Print("[ERROR] : Sending Auth to poly", err)
	}

	defer resp.Body.Close()

	//fmt.Println(resp.Body)
	if IsReaderAtLeast(resp.Body, 2000) {
		return []byte(`{"AuthResponse" : true}`)
	}
	return []byte(`{"AuthResponse" : false}`)
}

func IsReaderAtLeast(r io.Reader, size int) bool {
	tmp := make([]byte, size)
	
	// test //
	login := adapters.LoginBuilder{}
	login.GetParser(r).GetValue("value","input")
	//////////
	
	_, err := io.ReadAtLeast(r, tmp, size)
	if err == io.ErrUnexpectedEOF {
		return false
	}
	return true
}
