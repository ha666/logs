package logs

import (
	"os"
)

func init() {
	f := &PatternLogFormatter{
		Pattern:    "%w %t %F:%n - %m",
		WhenFormat: "2006-01-02 15:04:05.000",
	}
	dir, _ := os.Getwd()
	RegisterFormatter("pattern", f)
	_ = SetGlobalFormatter("pattern", dir)
}
