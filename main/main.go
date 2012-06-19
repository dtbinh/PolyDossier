// Le fichier principal qui contient l'entrée du fichier go.
package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"time"
	"studash/pages"
)

const kPolyHost = "https://www4.polymtl.ca"
const kLogFile  = "request"

func main() {
  file, _ := os.OpenFile(fmt.Sprintf("%s_%d%s", kLogFile, time.Now().Unix(), ".log"), os.O_WRONLY | os.O_CREATE, 0666 )
	log.SetOutput(file)
	
	http.HandleFunc("/", onHandleRequest)
	http.ListenAndServe(":http", nil)
	
	defer file.Close()
}

func onHandleRequest(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  doHandleRequest(w, r)
  log.Println(fmt.Sprintf("[%s] - %s", time.Since(start), r.URL.Path))
}

func doHandleRequest(w http.ResponseWriter, r *http.Request) {
  data := make([]byte, 0)
	
	switch r.URL.Path {
	  case "/" : data = DefaultPage()
		default : data = WeirdJSON()
	}
	
	w.Write(data)
}

func DefaultPage() []byte {
  file, err := os.Open("../src/studash/client/index.html")
	defer file.Close()
	if err != nil { fmt.Println("Opening default page : " + err.Error()) }
  
	data, err := ioutil.ReadAll(file)
	if err != nil { fmt.Println("Reading default page : " + err.Error()) }
	
	return data
}

func WeirdJSON() []byte {
  //return []byte("{name:lolsaure}")
	return pages.KAR
}
