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
  -c string
        config path to use (short)
  -config string
        config path to use
  -p uint
        port to use (short) (default 8888)
  -port uint
        port to use (default 8888)
  -r string
        root directory to use (short) (default "./")
  -root string
        root directory to use (default "./")
```

```
$ cd DOCUMENT_ROOT
$ goran
2017/11/07 23:22:29 Starting Goran HTTP Server
2017/11/07 23:22:29 Listen : http://127.0.0.1:8888
2017/11/07 23:22:29 RootDir: ./
2017/11/07 23:22:38 127.0.0.1:65208 GET /
```

### Configuration file
goran is configured with a simple [TOML](https://github.com/toml-lang/toml) file.

```
Port = 6000
Addr = "127.0.0.1"
rootDir = "/var/www/hoge"
```

## License

MIT License

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Copyright (c) @tkmru
