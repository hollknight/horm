package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewLogger 初始化 logger
func NewLogger() *zap.Logger {
	encoder := getEncoder()
	writer := getLogWriter()
	level := new(zapcore.Level)

	core := zapcore.NewCore(encoder, writer, level)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

// 配置 zap encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 配置 zap writer
func getLogWriter() zapcore.WriteSyncer {
	// 向控制台输出日志信息
	opts := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
	writeSyncer := zapcore.NewMultiWriteSyncer(opts...)

	return writeSyncer
}
