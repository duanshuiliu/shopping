package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

	config "shopping/admin/pkg/conf"
	helper "shopping/admin/pkg/helper"
)

var (
	Instance *zap.Logger

	defaultLogPath  string = "./log"
	defaultLogLevel string = "info"

	lastLoggerTime int
)

type logConfig struct {
	path  string
	level string
}

func NewLogger() *zap.Logger {
	cTime := time.Now()
	dTime := cTime.YearDay()

	if dTime == lastLoggerTime {
		if Instance != nil {
			return Instance
		}
	}

	logConf  := getLogConfig()
	filename := fmt.Sprintf("/%04d-%02d-%02d.log", cTime.Year(), cTime.Month(), cTime.Day())

	hook := lumberjack.Logger{
		Filename:   logConf.path+filename,    // 日志文件路径
		MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	level       := logConf.level
	atomicLevel := zap.NewAtomicLevel()

	switch (level) {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	case "dpanic":
		atomicLevel.SetLevel(zap.DPanicLevel)
	case "panic":
		atomicLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		atomicLevel.SetLevel(zap.FatalLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	// caller := zap.AddCaller()
	// 开启文件及行号
	// development := zap.Development()
	// 设置初始化字段
	fields := zap.Fields(zap.String("rid", helper.GetRequestID()))
	// 构造日志
	// logger := zap.New(core, caller, development, fields)
	Instance       = zap.New(core, fields)
	lastLoggerTime = dTime

	return Instance
}

func getLogConfig() *logConfig {
	logConfig := &logConfig {
		path : defaultLogPath,
		level: defaultLogLevel,
	}

	appConf, err := config.New("app")
	if err != nil { return logConfig }

	if filepath := appConf.GetString("log.path"); filepath != "" {
		logConfig.path = filepath
	}

	if level := appConf.GetString("log.level"); level != "" {
		logConfig.level = level
	}

	return logConfig
}