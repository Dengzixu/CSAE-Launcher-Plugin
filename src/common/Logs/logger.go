package Logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

var G *zap.SugaredLogger

func Init() {
	cd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logFileName := cd + "\\.csae\\log\\logs.log"

	hook := lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    1024,
		MaxAge:     3,
		MaxBackups: 7,
		LocalTime:  false,
		Compress:   true,
	}

	fw := zapcore.AddSync(&hook)
	cw := zapcore.Lock(os.Stdout)

	// 日志文件输出
	pec := zap.NewProductionEncoderConfig()
	pec.TimeKey = "time"
	pec.EncodeTime = zapcore.ISO8601TimeEncoder
	pec.EncodeDuration = zapcore.SecondsDurationEncoder
	pec.EncodeCaller = zapcore.ShortCallerEncoder
	pec.EncodeName = zapcore.FullNameEncoder

	// 控制台输出
	dec := zap.NewDevelopmentEncoderConfig()

	ct := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(pec), fw, zap.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(dec), cw, zap.DebugLevel),
	)

	logger := zap.New(ct, ).Sugar()

	G = logger
}
