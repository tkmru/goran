# goran
simple http server. 

## Installation

```
go get github.com/tkmru/goran
```

## Usage

```
$ ./goran -h
Usage:
  -a string
        address to use (short) (default "127.0.0.1")
  -address string
        address to use (default "127.0.0.1")
  -p int
        port to use (short) (default 8888)
  -port int
        port to use (default 8888)
```

```
$ cd DOCUMENT_ROOT
$ goran
2017/10/26 00:35:56 Start Goran HTTP Server
2017/10/26 00:35:56 http://127.0.0.1:8888
2017/10/26 00:36:15 127.0.0.1:57981 GET /
2017/10/26 00:36:15 127.0.0.1:57981 GET /favicon.ico
```
