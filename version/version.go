package version

import (
	"flag"
	"fmt"
	"os"
)

var versionFlag bool

var goVersion string
var gitCommit string
var branch    string
var buildTime string
var oSArch    string

func init() {
	flag.BoolVar(&versionFlag,"version", false, "print version info")
}

func PrintVersionOrContinue() {
	if versionFlag {
		fmt.Printf("Go version:  %v\n", goVersion)
		fmt.Printf("Git commit:  %v\n", gitCommit)
		fmt.Printf("branch:      %v\n", branch)
		fmt.Printf("Built:       %v\n", buildTime)
		fmt.Printf("OS/Arch:     %v\n", oSArch)

		os.Exit(0)
	}
}


