package cli

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"gopkg.in/natefinch/lumberjack.v2"
// )


// func NewLogger() *zap.Logger {
// 	logLevel := getLogLevel()
// 	encoderConfig := getEncoderConfig()
// 	writerSyncer := getWriterSyncer()
// 	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writerSyncer, logLevel)
// 	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
// 	return logger
// }

// func getLogLevel() zapcore.Level {
// 	logLevelStr := os.Getenv("LOG_LEVEL")
// 	level := zapcore.InfoLevel
// 	switch logLevelStr {
// 	case "debug":
// 		level = zapcore.DebugLevel
// 	case "info":
// 		level = zapcore.InfoLevel
// 	case "warn":
// 		level = zapcore.WarnLevel
// 	case "error":
// 		level = zapcore.ErrorLevel
// 	case "fatal":
// 		level = zapcore.FatalLevel
// 	default:
// 		fmt.Println("Invalid LOG_LEVEL, defaulting to INFO")
// 	}
// 	return level
// }

// func getEncoderConfig() zapcore.EncoderConfig {
// 	encodeConfig := zap.NewProductionEncoderConfig()
// 	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// 	encodeConfig.TimeKey = "Time"
// 	encodeConfig.LevelKey = "Level"
// 	encodeConfig.NameKey = "Name"
// 	encodeConfig.CallerKey = "Caller"
// 	encodeConfig.MessageKey = "Message"
// 	encodeConfig.StacktraceKey = "Stacktrace"
// 	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
// 	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
// 	return encodeConfig
// }

// func getWriterSyncer() zapcore.WriteSyncer {
// 	logFile := os.Getenv("LOG_FILE")
// 	if logFile == "" {
// 		logFile = "./log/log.txt"
// 	}
// 	maxSizeStr := os.Getenv("LOG_MAX_SIZE")
// 	maxSize := 100 // MB
// 	if maxSizeStr != "" {
// 		size, err := strconv.Atoi(maxSizeStr)
// 		if err == nil {
// 			maxSize = size
// 		} else {
// 			fmt.Println("Invalid LOG_MAX_SIZE, defaulting to 100MB")
// 		}
// 	}

// 	maxBackupsStr := os.Getenv("LOG_MAX_BACKUPS")
// 	maxBackups := 5
// 	if maxBackupsStr != "" {
// 		backups, err := strconv.Atoi(maxBackupsStr)
// 		if err == nil {
// 			maxBackups = backups
// 		} else {
// 			fmt.Println("Invalid LOG_MAX_BACKUPS, defaulting to 5")
// 		}
// 	}

// 	maxAgeStr := os.Getenv("LOG_MAX_AGE")
// 	maxAge := 7 // days
// 	if maxAgeStr != "" {
// 		age, err := strconv.Atoi(maxAgeStr)
// 		if err == nil {
// 			maxAge = age
// 		} else {
// 			fmt.Println("Invalid LOG_MAX_AGE, defaulting to 7 days")
// 		}
// 	}

// 	compressStr := os.Getenv("LOG_COMPRESS")
// 	compress := false
// 	if compressStr != "" {
// 		c, err := strconv.ParseBool(compressStr)
// 		if err == nil {
// 			compress = c
// 		} else {
// 			fmt.Println("Invalid LOG_COMPRESS, defaulting to false")
// 		}
// 	}

// 	lumberJackLogger := &lumberjack.Logger{
// 		Filename:   logFile,
// 		MaxSize:    maxSize,    // megabytes after which new file is created
// 		MaxBackups: maxBackups, // number of backups
// 		MaxAge:     maxAge,     // days
// 		Compress:   compress,
// 	}		
// 	return zapcore.AddSync(lumberJackLogger)
// }
