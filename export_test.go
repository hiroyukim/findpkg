package findpkg

func ExportSetFlagPkgs(s string) func() {
	org := flagPkg
	flagPkg = s
	return func() {
		flagPkg = org
	}
}
