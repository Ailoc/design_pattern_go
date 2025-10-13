package main

import (
	"fmt"
	"testing"
	"time"
)

// 装饰器模式是一种结构型模式，装饰器模式通过将对象嵌套在另一个对象中，为其添加额外的功能，而不改变其接口

// 比如一个基础的日志系统，只能打印普通的消息，在不修改原有日志系统代码的情况下，添加时间戳功能以及日志级别功能

type Logger1 interface {
	Log(message string)
}
type SimpleLogger struct{}

func (l *SimpleLogger) Log(message string) {
	println(message)
}

// 装饰器基类
type LoggerDecorator struct {
	wrapped Logger1
}

func (d *LoggerDecorator) Log(message string) {
	d.wrapped.Log(message)
}

// 添加时间戳功能的装饰器
type TimestampLogger struct {
	LoggerDecorator
}

func NewTimestampLogger(wrapped Logger1) *TimestampLogger {
	return &TimestampLogger{LoggerDecorator{wrapped: wrapped}}
}
func (d *TimestampLogger) Log(message string) {
	d.LoggerDecorator.Log(fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), message))
}

// 添加日志级别功能的装饰器
type LevelLogger struct {
	LoggerDecorator
	level string
}

func NewLevelLogger(wrapped Logger1, level string) *LevelLogger {
	return &LevelLogger{LoggerDecorator{wrapped: wrapped}, level}
}
func (d *LevelLogger) Log(message string) {
	d.LoggerDecorator.Log(fmt.Sprintf("[%s] %s", d.level, message))
}

func TestDecorator(t *testing.T) {
	// 创建一个简单的日志记录器
	var logger Logger1 = &SimpleLogger{}
	logger.Log("This is a simple log message.")
	// 使用时间戳装饰器
	logger = NewTimestampLogger(logger)
	logger.Log("This is a log message with timestamp.")
	// 使用日志级别装饰器
	logger = NewLevelLogger(logger, "INFO")
	logger.Log("This is an info log message with timestamp.")
}
