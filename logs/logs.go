package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zapcore.Field) {
	switch t := message.(type) {
	case error:
		log.Error(t.Error(), fields...)
	case string:
		log.Error(t, fields...)
	}

}
