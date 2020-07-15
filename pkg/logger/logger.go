package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger global logger util
var Logger *zap.Logger

// Init init logger
func Init() {
	hook := lumberjack.Logger{
		// 日志文件路径
		Filename: viper.GetString("log.logger_file"),
		// 每个日志文件保存的最大尺寸 单位：M
		MaxSize: viper.GetInt("log.log_rotate_size"),
		// 日志文件最多保存多少个备份
		MaxBackups: viper.GetInt("log.log_backup_count"),
		// 文件最多保存多少天
		MaxAge: viper.GetInt("log.log_rotate_date"),
		// 是否压缩
		Compress: true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "linenum",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		// 小写编码器
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		// ISO8601 UTC 时间格式
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		// 全路径编码器
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeName:   zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
}
