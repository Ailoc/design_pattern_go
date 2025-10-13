package main

import (
	"fmt"
	"testing"
)

// 备忘录模式是一种行为型设计模式，主要目的是在不破坏对象封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，以便以后恢复对象。
// 设计一个文本编辑器（Editor），可以保存编辑状态（如内容和光标位置），然后撤销（Undo）到之前的状态。这样，用户可以多次编辑并回滚。

type Memento struct { // 备忘录，存储状态
	content string
	cursor  int
}

type Editor struct { // 编辑器，需要保存状态
	content string
	cursor  int
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) SetContent(content string) {
	e.content = content
}
func (e *Editor) SetCursor(cursor int) {
	e.cursor = cursor
}

func (e *Editor) CreateMemento() *Memento { // 创建备忘录
	return &Memento{
		content: e.content,
		cursor:  e.cursor,
	}
}
func (e *Editor) Restore(m *Memento) { // 恢复备忘录
	e.content = m.content
	e.cursor = m.cursor
}
func (e *Editor) PrintState() { // 显示当前状态
	fmt.Printf("Content: %s, Cursor: %d\n", e.content, e.cursor)
}

// 管理备忘录栈
type History struct {
	mementos []*Memento
}

func NewHistory() *History {
	return &History{}
}
func (h *History) Push(m *Memento) { // 添加备忘录
	h.mementos = append(h.mementos, m)
}
func (h *History) Pop() *Memento { // 获取备忘录
	if len(h.mementos) > 0 {
		memento := h.mementos[len(h.mementos)-1]
		h.mementos = h.mementos[:len(h.mementos)-1]
		return memento
	}
	return nil
}
func TestMemento(t *testing.T) {
	editor := NewEditor()
	history := NewHistory()

	// 初始状态
	editor.SetContent("Hello")
	editor.SetCursor(5)
	editor.PrintState() // Current state: Content='Hello', Cursor=5

	// 保存状态1
	history.Push(editor.CreateMemento())

	// 修改状态
	editor.SetContent("Hello World")
	editor.SetCursor(11)
	editor.PrintState() // Current state: Content='Hello World', Cursor=11

	// 保存状态2
	history.Push(editor.CreateMemento())

	// 进一步修改
	editor.SetContent("Hi Universe")
	editor.SetCursor(10)
	editor.PrintState() // Current state: Content='Hi Universe', Cursor=10

	// 撤销到状态2
	if memento := history.Pop(); memento != nil {
		editor.Restore(memento)
		fmt.Println("After undo to state 2:")
		editor.PrintState() // Current state: Content='Hello World', Cursor=11
	}

	// 再次撤销到状态1
	if memento := history.Pop(); memento != nil {
		editor.Restore(memento)
		fmt.Println("After undo to state 1:")
		editor.PrintState() // Current state: Content='Hello', Cursor=5
	}
}
