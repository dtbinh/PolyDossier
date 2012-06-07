package main

import (
	"fmt"
        "net"
	"net/http"
	"net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	fmt.Printf("Starting server")
	l, _ := net.Listen("tcp", "127.0.0.1:9000")
	b := new(FastCGIServer)
	fcgi.Serve(l, b)
}
