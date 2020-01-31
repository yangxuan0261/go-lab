package syslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Access *zap.Logger
var Error *zap.SugaredLogger

func Init(afile, efile string, isDev bool) {
	Access = NewLogger(afile, isDev)
	Error = NewErrorLogger(efile, isDev)
}

func StdoutToggle() error {
	return nil
}

func NewLogger(logPath string, isDev bool) *zap.Logger {
	ops := []string{logPath}

	if isDev {
		ops = append(ops, "stdout")
	}

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      ops, // 指定 io 输出
		ErrorOutputPaths: []string{"stderr"},
	}
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:     "T",
		LevelKey:    "L",
		NameKey:     "N",
		MessageKey:  "M",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		//EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if isDev { // 开发模式下才开启 代码行号
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.CallerKey = "C"
	}

	var err error
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func NewDataLogger(logPath string, isDev bool) *zap.SugaredLogger {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "console",
		OutputPaths:      []string{logPath},
		ErrorOutputPaths: []string{"stderr"},
	}
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "M",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	if isDev {
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.CallerKey = "C"
	}

	var err error
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}

func NewErrorLogger(logPath string, isDev bool) *zap.SugaredLogger {
	if logPath == "" {
		logPath = "stdout"
	}
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "console",
		OutputPaths:      []string{logPath, "stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "T",
		MessageKey:     "M",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.CallerKey = "C"

	var err error
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
