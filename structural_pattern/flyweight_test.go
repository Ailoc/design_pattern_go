package main

import (
	"fmt"
	"testing"
)

// 享元模式是一种创建型模式，通过共享对象来减少内存占用和提高性能
// 享元模式的核心思想是区分外在状态和内在状态
// 内在状态：对象内部状态，对象内部状态是可共享的
// 外在状态：对象外部状态，对象外部状态是可改变的

// 文本编辑器系统，内在状态是候是字符，外在状态是坐标
type Flyweight interface {
	Display(extrinsicState *ExtrinsicState)
}

type ExtrinsicState struct {
	X, Y int
}

type Character struct {
	Char rune
}

func (c *Character) Display(extrinsicState *ExtrinsicState) {
	fmt.Printf("Display %c at (%d, %d)\n", c.Char, extrinsicState.X, extrinsicState.Y)
}

// 享元工厂，管理共享池
type FlyweightFactory struct {
	pool map[rune]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		pool: make(map[rune]Flyweight),
	}
}
func (f *FlyweightFactory) GetCharacter(char rune) Flyweight {
	if char, ok := f.pool[char]; ok {
		return char
	}
	// 创建新的字符，然后加入享元池
	newChar := &Character{Char: char}
	f.pool[char] = newChar
	return newChar
}
func TestFlyweight(t *testing.T) {
	// 创建享元工厂
	flyweightFactory := NewFlyweightFactory()
	// 创建字符
	c1 := flyweightFactory.GetCharacter('a')
	c2 := flyweightFactory.GetCharacter('a')
	c3 := flyweightFactory.GetCharacter('c')
	// 创建外在状态
	extrinsicState1 := &ExtrinsicState{X: 10, Y: 20}
	extrinsicState2 := &ExtrinsicState{X: 30, Y: 40}

	c1.Display(extrinsicState1)
	c2.Display(extrinsicState1)
	c3.Display(extrinsicState2)
	// 判断是否共享
	if c1 == c2 {
		fmt.Println("c1 and c2 are the same object")
	} else {
		fmt.Println("c1 and c2 are different objects")
	}
	// 输出：c1 and c2 are the same object
}
