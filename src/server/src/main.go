// Impléntation de base pour l'application PolyDossier.
// Créé par Eliott Mahou 
// et Maxime Lavigne <malavv> duguigne@gmail.com
package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const kPolyHost = "https://www4.polymtl.ca"

type RequestError struct {
	Method string
	ierror error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%s got %s", r.Method, r.ierror)
}

var (
	ErrMethod = errors.New("Method isn't supported")
)

func main() {
	http.HandleFunc("/", goHandler)
	http.ListenAndServe(":http", nil)
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	resp, err := goResponse(r)

	if err != nil {
		fmt.Println(r.Method, err)
		return
	}

	decoder := xml.NewDecoder(resp.Body)
	token, tErr := xml.Token(nil), error(nil)
	token, tErr = decoder.Token()
	for i := 0; tErr == nil; i++ {
		// if i == 1 {
		//	decoder.Skip()
		//} else {*/
			fmt.Println(token, tErr)
			token, tErr = decoder.Token()
		//}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(r.Method, err)
		return
	}

	w.Write(data)

	resp.Body.Close()

	elapsed := time.Since(t)

	if elapsed < 200*time.Millisecond {
		fmt.Println(time.Since(t).String() + "\t" + r.URL.Path)
	} else {
		fmt.Println(time.Since(t).String() + "\t" + r.URL.Path + "\t[Trop Long!!!]")
	}
}

func goResponse(r *http.Request) (*http.Response, error) {
	defaultClient := &http.Client{}

	// Si on arrive sur la page.
	if r.URL.Path == "/" {
		return defaultClient.Get(kPolyHost + "/poly/poly.html")
	}

	switch r.Method {
	case "GET":
		return defaultClient.Get(kPolyHost + r.URL.Path)
	case "POST":
		r.ParseForm()
		return defaultClient.PostForm(kPolyHost+r.URL.Path, r.Form)
	default:
		return nil, &RequestError{Method: r.Method, ierror: ErrMethod}
	}
	return nil, nil
}
