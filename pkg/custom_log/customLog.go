package customlog

import (
	"fmt"
	"log"
	"os"
	"path"
)

func Log() {
	LOGFILE := path.Join(os.TempDir(), "clog.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Log create...", "in", os.TempDir())
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile | log.LstdFlags
	iLog := log.New(f, "", LstdFlags)
	iLog.Println("Mastering Go 3rd editing")
	iLog.Println("Another log entry...")
}
