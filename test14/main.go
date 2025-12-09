package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
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
