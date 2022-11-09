package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func InitLogger(logLevel string) {
	pe := zap.NewProductionEncoderConfig()

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		log.Printf("No logging level or wrong value provided as \"%s\"\n", logLevel)
		level = zap.InfoLevel
	}
	log.Printf("Logging at %s level", level.String())
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	logger = zap.New(core).Sugar()
}

func Info(args ...interface{}) {
	logger.Info(args)
}
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}
func Debug(args ...interface{}) {
	logger.Debug(args)
}
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args)
}
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
