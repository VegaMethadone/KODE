package loger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func NewLogger() error {
	filename, err := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	level := zapcore.InfoLevel
	core := zapcore.NewCore(encoder, zapcore.AddSync(filename), level)
	logger := zap.New(core)
	logger.Info("Logger initialized")

	Logger = logger

	return nil
}
