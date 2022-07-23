package main

import (
	"github.com/lance6716/k4Const/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// thanks https://disaev.me/p/writing-useful-go-analysis-linter/#writing-a-simple-linter
	singlechecker.Main(analyzer.Analyzer)
}
