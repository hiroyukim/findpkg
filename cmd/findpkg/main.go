package main

import (
	"findpkg"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findpkg.Analyzer) }
