package zlog

import (
	"testing"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {

	Info("benchmark", zap.Any("config", ""))
}
