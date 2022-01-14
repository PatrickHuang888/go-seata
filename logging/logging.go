package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ZapConfig *zap.Config `yaml:"zapConfig"`
}

var (
	global *zap.SugaredLogger
)

func init() {
	config := &Config{}

	encConfig := zap.NewDevelopmentEncoderConfig()
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.StampMicro))
	}

	config.ZapConfig = &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.ZapConfig.Build(zap.AddCallerSkip(1))
	global = logger.Sugar()
}

func Debug(s string) {
	global.Debug(s)
}

func Debugf(template string, args ...interface{}) {
	global.Debugf(template, args...)
}

func Error(s string) {
	global.Error(s)
}

func Errorf(template string, args ...interface{}) {
	global.Errorf(template, args...)
}

func Info(s string) {
	global.Info(s)
}

func Infof(template string, args ...interface{}) {
	global.Infof(template, args...)
}

func Warnn(s string) {
	global.Warn(s)
}

func Warnf(template string, args ...interface{}) {
	global.Warnf(template, args...)
}
