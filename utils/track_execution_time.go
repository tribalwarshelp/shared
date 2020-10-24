package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

func TrackExecutionTime(log *logrus.Entry, fn func(), fnName string) func() {
	return func() {
		now := time.Now()
		log := log.WithField("fnName", fnName)
		log.Infof("%s: called", fnName)

		fn()

		log.
			WithField("executionTime", time.Since(now).String()).
			Infof("%s: finished executing", fnName)
	}
}
