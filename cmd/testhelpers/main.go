package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/testhelpers"
)

func main() { singlechecker.Main(testhelpers.Analyzer) }
