package logging

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	encConfig zapcore.EncoderConfig
	encoder   zapcore.Encoder

	stdout zapcore.WriteSyncer
	stderr zapcore.WriteSyncer

	level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	global Logging
)

func init() {
	encConfig = zap.NewDevelopmentEncoderConfig()
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.StampMicro))
	}

	encoder = zapcore.NewConsoleEncoder(encConfig)

	sout, closer, err := zap.Open("stdout")
	if err != nil {
		closer()
		panic(fmt.Errorf("failed to initialize logger: %w", err))
	}

	serr, closer, err := zap.Open("stderr")
	if err != nil {
		closer()
		panic(fmt.Errorf("failed to initialize logger: %w", err))
	}

	stdout = sout
	stderr = serr

	global = NewLogging(NewLogger())
}

func SetLevel(l zapcore.Level) {
	level.SetLevel(l)
}

func NewLogger(extraWs ...zapcore.WriteSyncer) *zap.Logger {
	wss := append([]zapcore.WriteSyncer{stdout}, extraWs...)
	ws := zapcore.NewMultiWriteSyncer(wss...)

	core := zapcore.NewCore(encoder, ws, level)
	return zap.New(core, zap.ErrorOutput(stderr), zap.AddStacktrace(zapcore.ErrorLevel))
}

type Logging struct {
	//logger  *zap.Logger
	sugared *zap.SugaredLogger
}

func NewLogging(logger *zap.Logger) Logging {
	return Logging{
		sugared: logger.Sugar(),
	}
}

func Error(err error) {
	global.sugared.Error(err)
}

func Errorf(template string, args ...interface{}) {
	global.sugared.Errorf(template, args...)
}

func Info(str string) {
	global.sugared.Info(str)
}

func Infof(template string, args ...interface{}) {
	global.sugared.Infof(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	global.sugared.Fatalw(msg, keysAndValues...)
}

func Warning(err interface{}) {
	global.sugared.Warn(err)
}

func Warningf(template string, args ...interface{}) {
	global.sugared.Warnf(template, args)
}
func Infow(msg string, keysAndValues ...interface{}) {
	global.sugared.Infow(msg, keysAndValues)
}
