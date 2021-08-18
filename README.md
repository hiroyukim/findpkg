# findpkg

[![pkg.go.dev][gopkg-badge]][gopkg]

`findpkg` findpkg find callings specified by findpkg.pkgs flag.

```go
package main

import "log"

func main() {
	log.Fatal("hoge")
}
```

```sh
$ go vet -vettool=`which findpkg` -findpkg.pkgs="log" main.go
# command-line-arguments
./main.go:3:8: log is ng
```
