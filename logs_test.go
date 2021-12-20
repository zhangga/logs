package logs

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	SetLogLevel(DebugLevel)
	Debug("this is a debug log.")
	Debugf("this is a debug log, Name: %v", "zaplog")
	CtxDebugf(context.Background(), "this is a debug log, Name: %v", "zaplog")
}

func TestLogWithConf(t *testing.T) {
	WithLoggerConf("./logger.default.yaml")
	Debug("this is a debug log.")
	Errorf("this is a error log, Name: %v", "zaplog")
	CtxDebugf(context.Background(), "this is a debug log, Name: %v", "zaplog")
}
