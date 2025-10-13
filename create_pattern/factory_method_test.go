package main

import (
	"fmt"
	"testing"
)

type Handler interface {
	Handle()
}

type HandlerFactory interface {
	CreateHandler() Handler
}

func HandleServe(factory HandlerFactory) { // 核心业务逻辑
	handler := factory.CreateHandler()
	handler.Handle()
}

type HandlerA struct{}

func (h *HandlerA) Handle() {
	fmt.Println("对象A处理Handle")
}

type HandlerB struct{}

func (h *HandlerB) Handle() {
	fmt.Println("对象B处理Handle")
}

type HandlerFactoryA struct{}

func (h *HandlerFactoryA) CreateHandler() Handler {
	return &HandlerA{}
}

type HandlerFactoryB struct{}

func (h *HandlerFactoryB) CreateHandler() Handler {
	return &HandlerB{}
}

func TestFactoryMethod(t *testing.T) {
	api := "A"
	var factory HandlerFactory
	switch api {
	case "A":
		factory = &HandlerFactoryA{}
	case "B":
		factory = &HandlerFactoryB{}
	}
	HandleServe(factory)
}
