package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(logLevel string, isDev bool) (*LoggerZap, error) {
	// Parse chuỗi log level (debug, info, error, ...) sang kiểu zapcore.Level
	level := parseLogLevel(logLevel)

	var encoder zapcore.Encoder // Bộ định dạng log (console hoặc JSON)

	// Chuẩn bị thư mục log cho cả dev và production
	logDir := filepath.Dir("./storage/logs/app.log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("không thể tạo thư mục log %s: %w", logDir, err)
	}

	// Cấu hình lumberjack logger cho file rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename: "./storage/logs/app.log", // File log lưu tại đây
		Compress: true,                     // Nén file log cũ
	}

	// Cấu hình khác nhau cho dev và production
	if isDev {
		// Môi trường dev: log ra cả console và file, định dạng dễ đọc, lưu nhiều debug info hơn
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		lumberjackLogger.MaxSize = 50      // File nhỏ hơn (50MB)
		lumberjackLogger.MaxBackups = 3    // Ít bản sao hơn
		lumberjackLogger.MaxAge = 3        // Giữ trong thời gian ngắn hơn (3 ngày)
	} else {
		// Môi trường production: log ra file với định dạng JSON
		encoder = newJSONEncoder()
		lumberjackLogger.MaxSize = 100     // File lớn hơn (100MB)
		lumberjackLogger.MaxBackups = 5    // Nhiều bản sao hơn
		lumberjackLogger.MaxAge = 7        // Giữ lâu hơn (7 ngày)
	}

	// Kiểm tra quyền ghi vào file
	file, err := os.OpenFile(lumberjackLogger.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file log %s: %w", lumberjackLogger.Filename, err)
	}
	file.Close()

	// Tạo một mảng cores để kết hợp nhiều đầu ra
	var cores []zapcore.Core

	// Thêm core cho file log (cả dev và production đều có)
	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), level))

	// Nếu là môi trường dev, thêm core log ra console
	if isDev {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level))
	}

	// Kết hợp các cores
	core := zapcore.NewTee(cores...)

	// Tạo logger với các option:
	// - AddCaller: thêm thông tin file:line vào log
	// - AddStacktrace: thêm stack trace nếu là error trở lên
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &LoggerZap{Logger: logger}, nil
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
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder // Log level viết hoa (DEBUG, INFO,...)
	cfg.EncodeCaller = zapcore.ShortCallerEncoder // Chỉ lấy file:line chứ không full path
	return zapcore.NewJSONEncoder(cfg)            // Trả về encoder dạng JSON
}