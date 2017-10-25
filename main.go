package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start Goran HTTP Server")
	fmt.Printf("http://%s", l.Addr())

	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
}
