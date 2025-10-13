package main

import (
	"log"
	"testing"

	"go.uber.org/zap"
)

type Logger interface { // 需要统一的日志接口
	Info(msg string)
	Error(msg string)
}

type StdLogger struct {
	logger *log.Logger
}

func (l *StdLogger) Print(msg string) {
	l.logger.Println(msg)
}

type StdLoggerAdapter struct { // 适配器
	logger *StdLogger
}

func (l *StdLoggerAdapter) Info(msg string) {
	l.logger.Print("Info:" + msg)
}
func (l *StdLoggerAdapter) Error(msg string) {
	l.logger.Print("Error:" + msg)
}

func NewStdLoggerAdapter(logger *StdLogger) *StdLoggerAdapter {
	return &StdLoggerAdapter{logger}
}

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) LogInfo(msg string) {
	l.logger.Info(msg)
}
func (l *ZapLogger) LogError(msg string) {
	l.logger.Error(msg)
}

type ZapLoggerAdapter struct { // 适配器
	logger *ZapLogger
}

func (l *ZapLoggerAdapter) Info(msg string) {
	l.logger.LogInfo("Info:" + msg)
}
func (l *ZapLoggerAdapter) Error(msg string) {
	l.logger.LogError("Error:" + msg)
}
func NewZapLoggerAdapter(logger *ZapLogger) *ZapLoggerAdapter {
	return &ZapLoggerAdapter{logger}
}

// 业务系统
type Application struct {
	logger Logger
}

func NewApplication(logger Logger) *Application {
	return &Application{logger}
}

func (a *Application) Run() {
	a.logger.Info("Application started")
	a.logger.Error("Application error")
}

func TestAdapter(t *testing.T) {
	StdLogger := &StdLogger{log.New(log.Writer(), "StdLogger: ", log.LstdFlags)}
	StdLoggerAdapter := NewStdLoggerAdapter(StdLogger)
	app := NewApplication(StdLoggerAdapter)
	app.Run()

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	ZapLoggerInstance := &ZapLogger{logger: zapLogger}
	ZapLoggerAdapter := NewZapLoggerAdapter(ZapLoggerInstance)
	app = NewApplication(ZapLoggerAdapter)
	app.Run()
}
