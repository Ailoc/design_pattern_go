package main

import (
	"fmt"
	"testing"
)

// 命令模式是一种行为型模式，命令模式将一个请求封装为一个对象，调用者不直接调用接收者的方法，而是通过命令对象间接执行

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	IsOn bool
}

func (l *Light) On() {
	l.IsOn = true
}
func (l *Light) Off() {
	l.IsOn = false
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
	fmt.Println("Light is ON")
}
func (c *LightOnCommand) Undo() {
	c.light.Off()
	fmt.Println("Light is OFF")
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
	fmt.Println("Light is OFF")
}
func (c *LightOffCommand) Undo() {
	c.light.On()
	fmt.Println("Light is ON")
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}
func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func TestCommand(t *testing.T) {

	light := &Light{}
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}
	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()
	if !light.IsOn {
		t.Error("Light should be ON")
	}
	remote.PressUndo()
	if light.IsOn {
		t.Error("Light should be OFF")
	}

	remote.SetCommand(lightOff)
	remote.PressButton()
	if light.IsOn {
		t.Error("Light should be OFF")
	}
	remote.PressUndo()
	if !light.IsOn {
		t.Error("Light should be ON")
	}
}
