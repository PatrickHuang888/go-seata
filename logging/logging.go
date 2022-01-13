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

	logger, _ := config.ZapConfig.Build()
	global = logger.Sugar()
}

func Debugf(template string, args ...interface{}) {
	global.Debugf(template, args)
}

func Debug(args ...interface{}) {
	global.Debug(args)
}

func Error(args interface{}) {
	global.Error(args)
}

func Errorf(template string, args ...interface{}) {
	global.Errorf(template, args...)
}

func Info(str string) {
	global.Info(str)
}

func Infof(template string, args ...interface{}) {
	global.Infof(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	global.Fatalw(msg, keysAndValues...)
}

func Warning(err interface{}) {
	global.Warn(err)
}

func Warningf(template string, args ...interface{}) {
	global.Warnf(template, args)
}
func Infow(msg string, keysAndValues ...interface{}) {
	global.Infow(msg, keysAndValues)
}
