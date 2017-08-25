package alarms

import (
	"testing"
	"time"
	"github.com/sirupsen/logrus"
)

func TestStart(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	Start()
	for {
		time.Sleep(100*time.Second)
	}
}

func TestMisc(t *testing.T) {
}

