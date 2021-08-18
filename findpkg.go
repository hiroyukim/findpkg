package findpkg

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = "findpkg is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findpkg",
	Doc:  doc,
	Run:  run,
}

var flagPkg string

func init() {
	Analyzer.Flags.StringVar(&flagPkg, "pkgs", "", "do not use this package")
}

func run(pass *analysis.Pass) (interface{}, error) {
	if flagPkg == "" {
		return nil, nil
	}

	for _, f := range pass.Files {
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if path == flagPkg {
				pass.Reportf(ip.Pos(), "%s is ng", path)
			}
		}
	}

	return nil, nil
}
