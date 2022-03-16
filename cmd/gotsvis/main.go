package main

import (
	"github.com/moromin/gotsvis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gotsvis.Analyzer) }
