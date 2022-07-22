package k4const

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"k4Const/pkg/analyzer"
)

func main() {
	// thanks https://disaev.me/p/writing-useful-go-analysis-linter/#writing-a-simple-linter
	singlechecker.Main(analyzer.Analyzer)
}
