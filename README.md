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

## License

MIT License

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Copyright (c) @tkmru
