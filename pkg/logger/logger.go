package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger 初始化日志
func InitLogger(mode string, logFile string) {
	var cfg zap.Config
	if mode == "dev" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	// 日志输出到文件
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("无法打开日志文件: %v", err)
		}
		ws := zapcore.AddSync(file)
		cfg.OutputPaths = []string{logFile}
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			ws,
			cfg.Level,
		)
		Log = zap.New(core, zap.AddCaller())
		return
	}

	// 默认输出到控制台
	var err error
	Log, err = cfg.Build()
	if err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
}

// 常用日志方法封装
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// 标准库log实现
func StdInfo(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

func StdWarn(format string, v ...interface{}) {
	log.Printf("[WARN] "+format, v...)
}

func StdDebug(format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}

func StdError(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
