package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(logLevel string, isDev bool) *LoggerZap {
	// Parse chuỗi log level (debug, info, error, ...) sang kiểu zapcore.Level
	level := parseLogLevel(logLevel)

	var encoder zapcore.Encoder         // Bộ định dạng log (console hoặc JSON)
	var writeSyncer zapcore.WriteSyncer

	if isDev {
		// Nếu là môi trường dev: log ra console, dễ đọc
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		// Nếu là production: log ra file, định dạng JSON, dùng lumberjack để quay vòng log
		encoder = newJSONEncoder()
		writeSyncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./storage/logs/app.log", // File log lưu tại đây
			MaxSize:    100,                      // Dung lượng tối đa 100MB
			MaxBackups: 5,                        // Tối đa 5 file backup
			MaxAge:     7,                        // Xoá sau 7 ngày
			Compress:   true,                     // Nén file log cũ
		})
	}

	// Kết hợp các thành phần để tạo thành core logger
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Tạo logger với các option:
	// - AddCaller: thêm thông tin file:line vào log
	// - AddStacktrace: thêm stack trace nếu là error trở lên
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &LoggerZap{Logger: logger}
}

// parseLogLevel chuyển đổi chuỗi log level (debug, info, ...) thành zapcore.Level
func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// newJSONEncoder cấu hình encoder dạng JSON cho production
func newJSONEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder       // Định dạng thời gian ISO8601
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder     // Log level viết hoa (DEBUG, INFO,...)
	cfg.EncodeCaller = zapcore.ShortCallerEncoder     // Chỉ lấy file:line chứ không full path
	return zapcore.NewJSONEncoder(cfg)                // Trả về encoder dạng JSON
}
