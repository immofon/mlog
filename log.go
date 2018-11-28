package mlog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.Level = logrus.InfoLevel

	fmter := new(logrus.JSONFormatter)
	fmter.DisableTimestamp = true

	logger.Formatter = fmter
	logger.ExitFunc = func(code int) {
		fmt.Println("exit", code)
	}
	logger.SetOutput(os.Stderr)
}

func TextMode() {
	logger.SetFormatter(new(logrus.TextFormatter))
}
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}
func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

func l(skip int) *logrus.Entry {
	pc, file, line, _ := runtime.Caller(skip)
	file = filepath.Base(file)
	return logger.
		WithField("§time", time.Now().Format("2006-01-02 15:04:05.9999999")).
		WithField("§file", file).
		WithField("§line", line).
		WithField("§func", runtime.FuncForPC(pc).Name())
}
func L() *logrus.Entry {
	return l(2)
}
func E(err error) *logrus.Entry {
	return l(2).WithField("$err", err)
}
