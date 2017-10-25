package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        http.ServeFile(w, r, "."+r.URL.Path)
    })

    log.Printf("Start Goran HTTP Server")
    addr := "127.0.0.1:8080"
    log.Printf("http://%s", addr)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal(err)
    }
}
