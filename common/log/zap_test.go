package log

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	SetLoggerLevel("debug")
	Debugf("this is %s log ...", "debugf")
	Info("this is info log ...")
	Infof("this is %s log ...", "infof")
	Infow("this is infow log",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("uid", 3),
		zap.Duration("backoff", time.Second),
	)
	Warnf("this is %s log ...", "warnf")
	Errorf("this is %s log ...", "errorf")
	Fatalf("this is %s log ...", "fatalf")
	DPanicf("this is %s log ...", "dPanicf")
}
