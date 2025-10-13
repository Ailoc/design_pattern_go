package main

import (
	"fmt"
	"sync"
	"testing"
)

type Logger struct {
	name string
}

func (l *Logger) Log(msg string) {
	println(l.name + ": " + msg)
}

var (
	logger *Logger
	once   sync.Once
)

func GetLogger() *Logger { // 懒汉式单例(饿汉式单例是在程序启动时就初始化，懒汉式是在第一次使用时初始化)
	once.Do(func() {
		logger = &Logger{name: "logger"}
		fmt.Println("logger init")
	})
	return logger
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			GetLogger().Log(fmt.Sprintf("hello world %d", i))
		}(i)
	}
	wg.Wait()

	logger1 := GetLogger()
	logger2 := GetLogger()
	fmt.Println(logger1 == logger2)
}
