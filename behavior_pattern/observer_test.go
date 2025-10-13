package main

import (
	"fmt"
	"testing"
)

// 也叫监听者模式或者事件订阅模式
// 观察者模式就是说当一个对象或者说被观察者的状态发生变化时，观察者会收到通知，并做出相应的处理。
// 在观察者模式中，每个观察者不会主动发起观察，而是被动的等待被观察者发出通知。

type Observer interface {
	Update(price float64)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	Notify() // 通知所有观察者
}

type StockManager struct { // 管理观察者, Subject的具体实现类
	price     float64
	observers []Observer
}

func NewStockManager(initialPrice float64) *StockManager {
	return &StockManager{price: initialPrice}
}

func (sm *StockManager) SetPrice(newPrice float64) {
	sm.price = newPrice
	fmt.Printf("股票价格变化为: %.2f\n", newPrice)
	sm.Notify()
}

func (sm *StockManager) RegisterObserver(observer Observer) {
	sm.observers = append(sm.observers, observer)
}
func (sm *StockManager) RemoveObserver(observer Observer) {
	for i, o := range sm.observers {
		if o == observer {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			break
		}
	}
}
func (sm *StockManager) Notify() {
	for _, observer := range sm.observers {
		observer.Update(sm.price)
	}
}

type Investor struct {
	name string
}

func NewInvestor(name string) *Investor {
	return &Investor{name: name}
}

func (i *Investor) Update(price float64) {
	fmt.Printf("%s 收到通知，股票价格变化为: %.2f\n", i.name, price)
	fmt.Println(i.decideAction(price))
}

// decideAction 简单逻辑：根据价格决定买/卖
func (i *Investor) decideAction(price float64) string {
	if price < 100 {
		return "买入！"
	} else if price > 150 {
		return "卖出！"
	}
	return "持有。"
}

func TestObserver(t *testing.T) {
	appleStock := NewStockManager(120.0)
	investor1 := NewInvestor("小王")
	investor2 := NewInvestor("小李")
	appleStock.RegisterObserver(investor1)
	appleStock.RegisterObserver(investor2)

	appleStock.SetPrice(110.0)
	appleStock.SetPrice(160.0)
	appleStock.SetPrice(90.0)
}
