// Connection au serveur et distribution élémentaire des tâches.
package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
	"os"
	"io/ioutil"
	"studash/errors"
	"strings"
	"studash/pages"
)

func onHandleRequest(w http.ResponseWriter, r *http.Request) {
    log.Println("[INFO] : Bagamon")
	start := time.Now()
	doHandleRequest(w, r)
	log.Println(fmt.Sprintf("[%s] - %s", time.Since(start), r.URL.Path))
}

func doHandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// w.Write(tools.ParseHTML(w,r))
		w.Write(DefaultPage())
	case "/js/base.js":
		w.Write(DefaultScript())
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

func DefaultScript() []byte {
	file, err := os.Open("../src/studash/client/js/base.js")
	defer file.Close()
	if err != nil {
		fmt.Println("Opening default script : " + err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Reading default script : " + err.Error())
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
	return []byte{}, errors.ErrUnimplemented
}

func doGet(path string) ([]byte, error) {
	c := pages.Credentials{"temp", "temp", "temp"}

	if path[0:3] == "/u/" {
		if strings.Contains(path[4:], "/") {
			endidx := strings.LastIndex(path[4:], "/")
			return pages.FetchInfos(c, path[5+endidx:]), nil
		} else {
			return pages.ListFunctions(c), nil
		}

	}
	return []byte{}, errors.ErrUnimplemented
}
func doPost(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}
func doPut(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}
func doDelete(path string) ([]byte, error) {
	return []byte{}, errors.ErrUnimplemented
}