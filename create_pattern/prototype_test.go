package main

// 原型模式区别于通过构造函数创建对象，原型模式通过复制对象创建对象。
// 通过Clone方法来创建对象
import (
	"fmt"
	"testing"
)

type Prototype interface {
	Clone() Prototype
	Description() string
}

type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Clone() Prototype {
	return &Circle{c.Radius, c.Color}
}
func (c *Circle) Description() string {
	return fmt.Sprintf("Circle with radius %d and color %s", c.Radius, c.Color)
}

func TestProtoType(t *testing.T) {
	circle := &Circle{5, "red"}
	clonedCircle := circle.Clone().(*Circle) // 类型断言回Circle
	fmt.Println(circle.Description())
	clonedCircle.Color = "blue"
	fmt.Println(clonedCircle.Description())
}
