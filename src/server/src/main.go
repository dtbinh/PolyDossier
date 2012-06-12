// Impléntation de base pour l'application PolyDossier.
// Créé par Eliott Mahou 
// et Maxime Lavigne <malavv> duguigne@gmail.com
package main

import (
	"fmt"
	"io/ioutil"
	"errors"
	"net/http"
)

type RequestError struct {
  Method string
  ierror error
}

func (r *RequestError) Error() string {
  return fmt.Sprintf("%s got %s", r.Method, r.ierror)
}

var (
  kPolyHost = "https://www4.polymtl.ca"
  ErrMethod = errors.New("Method isn't supported")  
)

func main() {
    http.HandleFunc("/", goHandler)
    http.ListenAndServe(":http", nil)
}

func goHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Currently serving : " + r.URL.Path)
  resp, err := goResponse(r)
  
  if err != nil {
    fmt.Println(r.Method, err)
    return
  }
  
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(r.Method, err)
    return
  }
  
  w.Write(data)
  
  resp.Body.Close()
}

func goResponse(r *http.Request) (*http.Response, error) {
  defaultClient := &http.Client{}
  
  // Si on arrive sur la page.
  if r.URL.Path == "/" {
    return defaultClient.Get(kPolyHost + "/poly/poly.html")
  }
  
  switch r.Method {
    case "GET"  :
      return defaultClient.Get(kPolyHost + r.URL.Path)
    case "POST" :
      return defaultClient.PostForm(kPolyHost + r.URL.Path, r.Form)
    default :
      return nil, &RequestError{Method: r.Method, ierror: ErrMethod}
  }
  return nil, nil
}