// Le fichier principal qui contient l'entrée du fichier go.
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"studash/pages"
	"time"
)

const kPolyHost = "https://www4.polymtl.ca"
const kLogFile = "request"

var kErrUnimplemented = errors.New("Unimplemented function")

func main() {
	file, _ := os.OpenFile(fmt.Sprintf("%s_%d%s", kLogFile, time.Now().Unix(), ".log"), os.O_WRONLY|os.O_CREATE, 0666)
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
	switch r.URL.Path {
	case "/":
		w.Write(DefaultPage())
	default:
		{
			data, err := doAction(r.Method, r.URL.Path)
			if err != nil {
				fmt.Println("Wrong method / path combination")
			}
			w.Write(data)
		}
	}
}

func DefaultPage() []byte {
	file, err := os.Open("../src/studash/client/index.html")
	defer file.Close()
	if err != nil {
		fmt.Println("Opening default page : " + err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Reading default page : " + err.Error())
	}

	return data
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
	return []byte{}, kErrUnimplemented
}

func doGet(path string) ([]byte, error) {
	if path[0:3] == "/u/" {
		return pages.ListFunctions(pages.Credentials{"malavv", "tt11tt62", "890821"}), nil
	}
	return []byte{}, kErrUnimplemented
}
func doPost(path string) ([]byte, error) {
	return []byte{}, kErrUnimplemented
}
func doPut(path string) ([]byte, error) {
	return []byte{}, kErrUnimplemented
}
func doDelete(path string) ([]byte, error) {
	return []byte{}, kErrUnimplemented
}
