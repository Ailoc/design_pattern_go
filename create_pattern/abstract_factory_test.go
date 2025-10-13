package main

import (
	"fmt"
	"testing"
)

type Button interface {
	Render()
}

type Checkbox interface {
	Check()
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WinFactory struct{}

func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}
func (w *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

func LoadGUI(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()
	button.Render()
	checkbox.Check()
}

type WinButton struct{}

func (w *WinButton) Render() {
	fmt.Println("Render WinButton")
}

type WinCheckbox struct{}

func (w *WinCheckbox) Check() {
	fmt.Println("Check WinCheckbox")
}

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}
func (m *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

type MacButton struct{}

func (m *MacButton) Render() {
	fmt.Println("Render MacButton")
}

type MacCheckbox struct{}

func (m *MacCheckbox) Check() {
	fmt.Println("Check MacCheckbox")
}

func TestAbstractFactory(t *testing.T) {
	sysos := "win"
	switch sysos {
	case "win":
		LoadGUI(&WinFactory{})
	case "mac":
		LoadGUI(&MacFactory{})
	}
}
