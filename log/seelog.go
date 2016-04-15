package main

import (
	"github.com/cihub/seelog"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	defer seelog.Flush()

	seelog.Info("hello")
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	seelog.Info(path)

	loger, err := seelog.LoggerFromConfigAsFile("/seelog.xml")
	if err != nil {
		seelog.Error("err ", err)
		return
	}
	seelog.ReplaceLogger(loger)

	u := int(1)
	seelog.Info(u)

	seelog.Info("info")
	seelog.Error("error")
}