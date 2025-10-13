package main

import (
	"fmt"
	"testing"
)

// 模板方法定义了一个算法的骨架，将算法的某些步骤延迟到子类中。

// BrewStrategy 接口：钩子（冲泡步骤）
type BrewStrategy interface {
	Brew() string
}

// Beverage 模板：冲泡骨架
type Beverage struct {
	strategy BrewStrategy
}

func NewBeverage(strategy BrewStrategy) *Beverage {
	return &Beverage{strategy: strategy}
}

func (b *Beverage) Prepare() {
	fmt.Println("1. 加热水...")
	fmt.Println("2.", b.strategy.Brew()) // 钩子
	fmt.Println("3. 倒入杯子。")
}

// Tea 茶的具体实现
type Tea struct{}

func (t *Tea) Brew() string {
	return "浸泡茶叶"
}

// Coffee 咖啡的具体实现
type Coffee struct{}

func (c *Coffee) Brew() string {
	return "冲泡咖啡粉"
}

func TestTemplateMethod(t *testing.T) {
	// 冲泡茶
	tea := NewBeverage(&Tea{})
	tea.Prepare()
	// 输出：
	// 1. 加热水...
	// 2. 浸泡茶叶
	// 3. 倒入杯子。

	fmt.Println("---")

	// 冲泡咖啡
	coffee := NewBeverage(&Coffee{})
	coffee.Prepare()
	// 输出：
	// 1. 加热水...
	// 2. 冲泡咖啡粉
	// 3. 倒入杯子。
}
