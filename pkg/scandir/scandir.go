package scandir

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func ScanDirectory(path string) error {
	fmt.Println(path)

	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		filePath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}

	return nil
}
