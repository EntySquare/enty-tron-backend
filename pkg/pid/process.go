package pid

import (
	"os"
	"path/filepath"
)

func PassOrPanic() {
	alias := getCurProcessName()
	pid, err := GetPid(alias)
	if err != nil {
		panic("os error get pid")
	}
	if pid != "" {
		panic("started pid :" + pid + " as process " + alias)
	}
}

func getCurProcessName() string {
	path, _ := os.Executable()
	_, exec := filepath.Split(path)
	return exec
}
