package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    //"github.com/BurntSushi/toml"
)

var port int
var addr string
var configPath string

func usage() {
    fmt.Println("Usage:")
    flag.PrintDefaults()
    os.Exit(0)
}

func init() {
    const defaultPort = 8888
    flag.IntVar(&port, "port", defaultPort, "port to use")
    flag.IntVar(&port, "p", defaultPort, "port to use (short)")

    const defaultAddr = "127.0.0.1"
    flag.StringVar(&addr, "address", defaultAddr, "address to use")
    flag.StringVar(&addr, "a", defaultAddr, "address to use (short)")

    const defaultConfig = ""
    flag.StringVar(&configPath, "config", defaultConfig, "address to use")
    flag.StringVar(&configPath, "c", defaultConfig, "address to use (short)")
    flag.Usage = func() { usage() }
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func main() {
    flag.Parse()
    if configPath != "" {
        if !isExist(configPath) {
            log.Panic("Config file not found.")
        }
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        w.Header().Add("Cache-Control", "no-store")
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
