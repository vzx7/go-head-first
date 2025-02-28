package main

import (
	"github.com/vzx7/go-head-first/pkg/scandir"
)

func main() {
	defer scandir.ReportPanic()
	scandir.ScanDirectory("/")
}
