package main

import (
	"fmt"
	"io/ioutil"
	"errors"
	"net/http"
	//"net/http/fcgi"
	"crypto/tls"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
w.Write([]byte("array"))
	fmt.Println("Currently serving : " + req.URL.Path)
	return
	tr := &http.Transport{
	  TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	path := ""
	post := false

	switch req.URL.Path {
	  case "/" : path ="/poly/poly.html"
          case "/servlet/ValidationServlet": {
	    req.ParseForm()
	    post = true
	    path = req.URL.Path
	  }
          default : path = req.URL.Path
	}

	resp := &http.Response{}
	err := errors.New("l")

	if !post {
          resp, err = client.Get("https://www4.polymtl.ca" + path)	
	} else {
	  resp, err = client.PostForm("https://www4.polymtl.ca" + path, req.Form)
	}


	if err != nil {
	  w.Write([]byte(err.Error()))
	}
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	w.Write(body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Currently serving : " + r.URL.Path)

	tr := &http.Transport{
	  TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	path := ""
	post := false

	switch r.URL.Path {
	  case "/" : path ="/poly/poly.html"
          case "/servlet/ValidationServlet": {
	    r.ParseForm()
	    post = true
	    path = r.URL.Path
	  }
          default : path = r.URL.Path
	}

	resp := &http.Response{}
	err := errors.New("l")

	if !post {
          resp, err = client.Get("https://www4.polymtl.ca" + path)	
	} else {
	  resp, err = client.PostForm("https://www4.polymtl.ca" + path, r.Form)
	}


	if err != nil {
      w.Write([]byte(err.Error()))	
	}
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	w.Write(body)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}
