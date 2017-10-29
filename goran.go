package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/BurntSushi/toml"
)

var port uint
var addr string
var configPath string

func usage() {
    fmt.Println("Usage:")
    flag.PrintDefaults()
    os.Exit(0)
}

func init() {
    const defaultPort = 8888
    flag.UintVar(&port, "port", defaultPort, "port to use")
    flag.UintVar(&port, "p", defaultPort, "port to use (short)")

    const defaultAddr = "127.0.0.1"
    flag.StringVar(&addr, "address", defaultAddr, "address to use")
    flag.StringVar(&addr, "a", defaultAddr, "address to use (short)")

    const defaultConfig = ""
    flag.StringVar(&configPath, "config", defaultConfig, "config path to use")
    flag.StringVar(&configPath, "c", defaultConfig, "config path to use (short)")
    flag.Usage = func() { usage() }
}


type Config struct {
    Port uint
    Addr string
}

func main() {
    flag.Parse()
    var config Config
    if configPath != "" {
        if _, err := toml.DecodeFile(configPath, &config); err != nil {
            log.Panic(err)
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
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}
