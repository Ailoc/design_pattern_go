package main

import "testing"

// 策略模式定义了一些列的算法，将每个算法封装起来，并使他们可以相互替换。让算法的变化独立于使用它的客户端。

type PaymentStrategy interface {
	Pay(amount float64)
}

type AliPay struct{}

func (a *AliPay) Pay(amount float64) {
	println("使用支付宝支付：", amount)
}

type WeChatPay struct{}

func (w *WeChatPay) Pay(amount float64) {
	println("使用微信支付：", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}
func (p *PaymentContext) Pay(amount float64) {
	if p.strategy != nil {
		p.strategy.Pay(amount)
	} else {
		panic("请先设置支付方式")
	}
}
func TestStrategy(t *testing.T) {
	paymentContext := &PaymentContext{}
	paymentContext.SetStrategy(&AliPay{})
	paymentContext.Pay(100)
	paymentContext.SetStrategy(&WeChatPay{})
	paymentContext.Pay(200)
}
