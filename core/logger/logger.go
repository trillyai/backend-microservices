package logger

import (
	"fmt"
	"os"

	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

var LogLevel string

type Logger struct {
	Logger        *zap.Logger
	ComponentName string
}

func NewLogger(componentName string) *Logger {

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "ts",
			LevelKey:    "level",
			MessageKey:  "msg",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
	}

	return &Logger{Logger: logger, ComponentName: componentName}

}

// Debug logs a message with additional fields at DebugLevel.
func (l *Logger) Debug(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Debugw(formatMessage(l.ComponentName, message), keyValues...)
}

// Info logs a message with additional fields at InfoLevel.
func (l *Logger) Info(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Infow(formatMessage(l.ComponentName, message), keyValues...)
}

// Warn logs a message with additional fields at WarnLevel.
func (l *Logger) Warn(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Warnw(formatMessage(l.ComponentName, message), keyValues...)
}

// Error logs a message with additional fields at ErrorLevel.
func (l *Logger) Error(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Errorw(formatMessage(l.ComponentName, message), keyValues...)
}

// Panic logs a message with additional fields at PanicLevel.
func (l *Logger) Panic(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Panicw(formatMessage(l.ComponentName, message), keyValues...)
}

// Fatal logs a message with additional fields at FatalLevel.
func (l *Logger) Fatal(message string, keyValues ...interface{}) {
	l.Logger.Sugar().Fatalw(formatMessage(l.ComponentName, message), keyValues...)
}

func formatMessage(componentName, message string) string {
	return fmt.Sprintf("%s\t%s", componentName, message)
}
