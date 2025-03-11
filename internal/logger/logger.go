package logger

import (
	"github.com/quietdevil/Platform_common/pkg/closer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
	closer.Add(globalLogger.Sync)
}

func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}
