package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"errors"
	"net/http"
	"net/http/fcgi"
	"crypto/tls"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Currently serving : " + req.URL.Path)

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

func main() {
	fmt.Println("Starting server")
	l, _ := net.Listen("tcp", "127.0.0.1:9000")
	b := new(FastCGIServer)
	fcgi.Serve(l, b)
}
