package main

import (
	"fmt"
	"testing"
)

// 桥接模式就是将抽象部分与实现部分分离，使它们都可以独立地变化。
// 抽象部分定义了对象的行为，而实现部分则定义了具体如何执行该行为。
// 在面向对象的设计方法中，等同于将继承改为了组合。

// BankAPI 接口：实现部分（桥），定义银行支付操作
type BankAPI interface {
	ProcessPayment(amount float64) string // 处理支付
}

// VisaBank 具体实现：Visa银行API
type VisaBank struct{}

func (vb *VisaBank) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Visa银行处理支付: %.2f 元，扣款成功（国际费率）", amount)
}

// MasterCardBank 另一个具体实现：MasterCard银行API
type MasterCardBank struct{}

func (mcb *MasterCardBank) ProcessPayment(amount float64) string {
	return fmt.Sprintf("MasterCard银行处理支付: %.2f 元，扣款成功（本地费率）", amount)
}

// PaymentMethod 接口：抽象部分，持有BankAPI的桥
type PaymentMethod interface {
	Pay(amount float64) string // 执行支付
}

// CreditCard 扩展抽象：信用卡支付-----------------------组合API接口-------------------------
type CreditCard struct {
	bank BankAPI // 桥：组合BankAPI接口
}

func NewCreditCard(bank BankAPI) *CreditCard {
	return &CreditCard{bank: bank}
}

func (cc *CreditCard) Pay(amount float64) string {
	// 信用卡特有逻辑：加分期检查
	if amount > 1000 {
		return "信用卡支付: 分期选项可用。 " + cc.bank.ProcessPayment(amount)
	}
	return "信用卡支付: 一次性付款。 " + cc.bank.ProcessPayment(amount)
}

// MobilePay 另一个扩展抽象：移动支付
type MobilePay struct {
	bank BankAPI // 桥：组合BankAPI接口 ------------------------组合API接口-------------------------
}

func NewMobilePay(bank BankAPI) *MobilePay {
	return &MobilePay{bank: bank}
}

func (mp *MobilePay) Pay(amount float64) string {
	// 移动支付特有逻辑：加扫码验证
	return "移动支付: 扫码验证通过。 " + mp.bank.ProcessPayment(amount)
}

func TestBridge(t *testing.T) {
	// 创建实现（银行API）
	visa := &VisaBank{}
	mastercard := &MasterCardBank{}

	// 创建抽象（支付方式），注入桥（bank）
	creditVisa := NewCreditCard(visa)
	creditMaster := NewCreditCard(mastercard)
	mobileVisa := NewMobilePay(visa)
	mobileMaster := NewMobilePay(mastercard)

	// 执行支付
	fmt.Println(creditVisa.Pay(500.0))
	fmt.Println(creditMaster.Pay(1500.0))
	fmt.Println(mobileVisa.Pay(200.0))
	fmt.Println(mobileMaster.Pay(800.0))
}
