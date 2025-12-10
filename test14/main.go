package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func initLogger() {
	write, _ := os.OpenFile("test.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logrus.SetOutput(io.MultiWriter(os.Stdout, write))
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
}
func main() {
	initLogger()
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")
	logrus.SetReportCaller(true)
	logrus.Info("info msg with caller")
	log := logrus.WithFields(logrus.Fields{
		"user_id": 681200,
		"action":  "test",
		"time":    "2025-12-9",
	})
	log.Info("msg")

}
