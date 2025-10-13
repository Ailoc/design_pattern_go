package main

import (
	"fmt"
	"testing"
)

type Handler interface {
	SetNext(Handler)
	Handle(level string, message string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) { b.next = h }

type InfoHandler struct {
	BaseHandler
}

func (i *InfoHandler) Handle(level string, message string) {
	if level == "INFO" {
		fmt.Println("InfoHandler: " + message)
	} else if i.next != nil {
		i.next.Handle(level, message)
	} else {
		fmt.Println("No handler found for level: " + level)
	}
}

type WarnHandler struct {
	BaseHandler
}

func (w *WarnHandler) Handle(level string, message string) {
	if level == "WARNING" {
		fmt.Println("WarnHandler: " + message)
	} else if w.next != nil {
		w.next.Handle(level, message)
	} else {
		fmt.Println("No handler found for level: " + level)
	}
}

type ErrorHandler struct {
	BaseHandler
}

func (e *ErrorHandler) Handle(level string, message string) {
	if level == "ERROR" {
		fmt.Println("ErrorHandler: " + message)
	} else if e.next != nil {
		e.next.Handle(level, message)
	} else {
		fmt.Println("No handler found for level: " + level)
	}
}

func TestChainOfResponsibility(t *testing.T) {
	InfoHandler := &InfoHandler{}
	WarnHandler := &WarnHandler{}
	ErrorHandler := &ErrorHandler{}

	InfoHandler.SetNext(WarnHandler)
	WarnHandler.SetNext(ErrorHandler)

	// 测试不同级别的日志
	fmt.Println("处理 Info 日志：")
	InfoHandler.Handle("INFO", "系统启动")

	fmt.Println("\n处理 Warning 日志：")
	InfoHandler.Handle("WARNING", "内存使用率高")

	fmt.Println("\n处理 Error 日志：")
	InfoHandler.Handle("ERROR", "数据库连接失败")

	fmt.Println("\n处理未知级别日志：")
	InfoHandler.Handle("DEBUG", "调试信息")
}
