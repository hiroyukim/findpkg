package findpkg_test

import (
	"testing"

	"findpkg"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	defer findpkg.ExportSetFlagPkgs("fmt")()
	analysistest.Run(t, testdata, findpkg.Analyzer, "a")
}
