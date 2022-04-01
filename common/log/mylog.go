package log

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() error {
	Log = logrus.New()
	fileName := "./log/" + time.Now().Format("2006-01-02") + ".log"
	writer, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	Log.SetLevel(logrus.DebugLevel)
	Log.SetOutput(io.Writer(writer))
	//Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetFormatter(&logrus.TextFormatter{DisableColors: false, TimestampFormat: "2006-01-02 15:04:05"})
	return nil
}
