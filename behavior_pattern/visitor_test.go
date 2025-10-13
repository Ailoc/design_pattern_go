package main

import (
	"fmt"
	"testing"
)

// 访问者模式是一种行为型设计模式，用于将对象与操作算法进行分离。
// 对象包含一个Accept方法，该方法接收一个Visitor对象作为参数，并调用Visitor对象的Visit方法

type Shape interface {
	Accept(v Visitor)
}
type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	fmt.Println("Circle area:", c.radius*c.radius*3.14)
}
func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	fmt.Println("Rectangle area:", r.width*r.height)
}

func TestVisitor(t *testing.T) {
	circle := &Circle{radius: 5}
	rectangle := &Rectangle{width: 4, height: 6}

	areaCalculator := &AreaCalculator{}

	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
}
