package logger

import (
	"battles/internal/utils/env"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"sync"
)

var logg *logrus.Logger = nil
var once sync.Once

func Get() *logrus.Logger {
	once.Do(func() {
		logg = &logrus.Logger{
			Out: os.Stderr,
			Formatter: &prefixed.TextFormatter{
				DisableColors:   true,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
				ForceFormatting: true,
			},
		}
		env.InitEnv()
		switch os.Getenv("logging") {
		case "warning":
			logg.SetLevel(logrus.WarnLevel)
		case "notice":
			logg.SetLevel(logrus.InfoLevel)
		case "debug":
			logg.SetLevel(logrus.DebugLevel)
		default:
			logg.SetLevel(logrus.InfoLevel)
		}
		logg.Infof("log level: %s", logg.Level.String())
	})
	return logg
}
