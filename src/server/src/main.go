package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/fcgi"
	"crypto/tls"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Nous avons besoin de specifier que nous ne voulons pas
	// réellement vérifier le Certifica de Poly.
	tr := &http.Transport{
	  TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://www4.polymtl.ca/poly/poly.html")

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
