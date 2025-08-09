package config

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	var (
		logger     *zap.Logger
		onceLogger sync.Once
	)

	onceLogger.Do(func() {
		logFile := &lumberjack.Logger{
			Filename:   viper.GetString("LOG_FILE"),
			MaxSize:    viper.GetInt("LOG_MAX_SIZE"),
			MaxAge:     viper.GetInt("LOG_MAX_AGE"),
			MaxBackups: viper.GetInt("LOG_MAX_BACKUPS"),
			Compress:   viper.GetBool("LOG_COMPRESS"),
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		core := zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(logFile), zapcore.InfoLevel),
		)

		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
	return logger
}
