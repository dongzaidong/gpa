package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	File string `json:"file"`
}

var Log *logrus.Logger

var debug = false
var defaultFile = "log.out"

func init() {
	Log = logrus.New()

	Log.SetLevel(logrus.InfoLevel)

	if debug {
		Log.SetLevel(logrus.DebugLevel)
	}

	f, err := os.OpenFile(defaultFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		f.Close()
		return
	}
	// 输出到终端与文件中
	o := io.MultiWriter(os.Stderr, f)
	Log.Out = o

}
func LogSettings() {

}

func Debug(args ...interface{}) {
	// entry := Log.WithField("pkg", "ssss")
	Log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}
