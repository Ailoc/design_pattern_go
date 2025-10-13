package main

import (
	"fmt"
	"testing"
)

// 中介者模式下，每个用户之间并不直接交换消息，而是通过中介者进行转发，由中介者进行协调
// 例如：聊天室，用户之间不能直接相互发送消息，只能通过聊天室进行转发

// 用户行为接口
type UserAction interface {
	Send(msg string)
	Receive(msg string)
	GetName() string // 获取用户名称
}

// 中介者接口
type Mediator interface {
	sendMessage(sender UserAction, msg string)
}

// 聊天室结构体，作为中介者
type ChatRoom struct {
	users []UserAction
}

func (c *ChatRoom) Register(u UserAction) {
	c.users = append(c.users, u)
	fmt.Println(u.GetName(), "加入聊天室")
}
func (c *ChatRoom) sendMessage(sender UserAction, msg string) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(msg)
		}
	}
}

// 聊天室用户结构体
type RoomUser struct {
	name     string
	mediator Mediator
}

func NewUser(name string, mediator Mediator) *RoomUser { // 创建用户
	return &RoomUser{name: name, mediator: mediator}
}

func (ru *RoomUser) Send(msg string) {
	ru.mediator.sendMessage(ru, msg)
}
func (ru *RoomUser) Receive(msg string) {
	fmt.Println(ru.name, "收到消息:", msg)
}
func (ru *RoomUser) GetName() string {
	return ru.name
}

func TestMediator(t *testing.T) {

	// 创建聊天室（中介者）
	chatRoom := &ChatRoom{}

	// 创建用户，并注册到聊天室
	alice := NewUser("Alice", chatRoom)
	bob := NewUser("Bob", chatRoom)
	charlie := NewUser("Charlie", chatRoom)

	chatRoom.Register(alice)
	chatRoom.Register(bob)
	chatRoom.Register(charlie)

	// 用户发送消息（通过中介者）
	alice.Send("大家好！")
	bob.Send("嗨，Alice！")
	charlie.Send("你们好啊。")

}
