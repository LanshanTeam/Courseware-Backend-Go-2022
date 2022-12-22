package main

import (
	"net"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website")))
	listener, err := net.Listen("tcp", "0.0.0.0:8077")
	if err != nil {
		panic(err)
	}
	panic(http.Serve(listener, http.DefaultServeMux))
}
