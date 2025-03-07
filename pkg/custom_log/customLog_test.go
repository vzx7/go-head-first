package customlog

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	LOGFILE := path.Join(os.TempDir(), "clog.log")
	Log()
	_, err := os.Stat(LOGFILE)
	if err != nil {
		t.Error("Error!", err)
	}

	fContent, err := os.ReadFile(LOGFILE)
	if err != nil {
		t.Error("Error!", err)
	}
	s := string(fContent)
	l := len(strings.Split(s, ""))
	if l == 0 {
		t.Error("Error, file empty.")
	}
	fmt.Println(s)
}
