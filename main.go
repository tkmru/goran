package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int
var addr string

func init() {
	const defaultPort = 8888
	flag.IntVar(&port, "port", defaultPort, "port to use")
	flag.IntVar(&port, "p", defaultPort, "port to use (short)")
	const defaultAddr = "127.0.0.1"
	flag.StringVar(&addr, "address", defaultAddr, "address to use")
	flag.StringVar(&addr, "a", defaultAddr, "address to use (short)")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	log.Printf("Start Goran HTTP Server")
	addr = fmt.Sprintf("%s:%d", addr, port)
	log.Printf("http://%s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
