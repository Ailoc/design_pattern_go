package main

import (
	"fmt"
	"testing"
)

// 状态模式是一种行为型模式，状态模式允许对象在其内部状态改变时改变其行为。
// 它将状态封装成独立的类，并将行为委托给当前状态对象，从而避免在主类中使用大量条件语句（if-else/switch）来处理不同状态。
// 这种模式就像给对象一个“状态机”，每个状态类定义了特定状态下的行为，当状态切换时，对象的行为随之改变。

// 场景：自动售货机（Vending Machine）。状态：NoCoin（无币）、HasCoin（有币）、Sold（已售）。
// 事件：InsertCoin（投币）、SelectItem（选购）。实际如电商或游戏状态机。
// 状态接口
type State interface {
	InsertCoin(m *VendingMachine) State
	SelectItem(m *VendingMachine) State
	String() string
}

// 上下文：售货机
type VendingMachine struct {
	state State
	item  string
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &NoCoinState{}}
}

func (m *VendingMachine) InsertCoin() {
	m.state = m.state.InsertCoin(m)
	fmt.Printf("State: %s\n", m.state)
}

func (m *VendingMachine) SelectItem() {
	m.state = m.state.SelectItem(m)
	fmt.Printf("State: %s, Item: %s\n", m.state, m.item)
}

// 无币状态
type NoCoinState struct{}

func (s *NoCoinState) InsertCoin(m *VendingMachine) State {
	fmt.Println("  -> Coin inserted")
	return &HasCoinState{}
}

func (s *NoCoinState) SelectItem(m *VendingMachine) State {
	fmt.Println("  -> No coin")
	return s
}

func (s *NoCoinState) String() string { return "NoCoin" }

// 有币状态
type HasCoinState struct{}

func (s *HasCoinState) InsertCoin(m *VendingMachine) State {
	fmt.Println("  -> Already has coin")
	return s
}

func (s *HasCoinState) SelectItem(m *VendingMachine) State {
	fmt.Println("  -> Dispensing item")
	m.item = "cola"
	return &SoldState{}
}

func (s *HasCoinState) String() string { return "HasCoin" }

// 已售状态
type SoldState struct{}

func (s *SoldState) InsertCoin(m *VendingMachine) State {
	fmt.Println("  -> Already sold")
	return s
}

func (s *SoldState) SelectItem(m *VendingMachine) State {
	fmt.Println("  -> Sold out")
	return &NoCoinState{}
}

func (s *SoldState) String() string { return "Sold" }
func TestState(t *testing.T) {
	vm := NewVendingMachine()
	vm.InsertCoin() // NoCoin -> HasCoin
	vm.SelectItem() // HasCoin -> Sold
	vm.InsertCoin() // Sold -> Sold (无效)
	vm.SelectItem() // Sold -> NoCoin
}
