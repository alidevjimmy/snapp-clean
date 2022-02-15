package zap

import (
	"github.com/alidevjimmy/snapp-clean/internal/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type log struct {
	zap *zap.SugaredLogger
}

func New(writer io.Writer, level zapcore.Level) logger.Logger {
	ws := zapcore.AddSync(writer)
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encodeConfig)
	core := zapcore.NewCore(enc, ws, level)

	z := zap.New(core)
	sugarLogger := z.Sugar()
	return &log{
		zap: sugarLogger,
	}
}

func (l *log) Error(msg string, kv ...interface{}) {
	l.zap.Errorw(msg, kv)
}

func (l *log) Warn(msg string, kv ...interface{}) {
	l.zap.Warnw(msg, kv)
}

func (l *log) Info(msg string, kv ...interface{}) {
	l.zap.Infow(msg, kv)
}

func (l *log) Sync() {
	l.zap.Sync()
}



