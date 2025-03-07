package customlog

import (
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
	l := len(strings.Split(string(fContent), ""))
	if l == 0 {
		t.Error("Error, file empty.")
	}
}
