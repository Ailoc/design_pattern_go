package main

import (
	"fmt"
	"testing"
)

// 子系统：DVD播放器
type DVDPlayer struct {
	currentMovie string
}

func (d *DVDPlayer) on() {
	fmt.Println("DVD Player is on")
}

func (d *DVDPlayer) play(movie string) {
	d.currentMovie = movie
	fmt.Println("Playing movie:", movie)
}

func (d *DVDPlayer) off() {
	fmt.Println("DVD Player is off")
}

// 子系统：投影仪
type Projector struct {
	input string
}

func (p *Projector) on() {
	fmt.Println("Projector is on")
}

func (p *Projector) setInput(input string) {
	p.input = input
	fmt.Println("Projector input set to:", input)
}

func (p *Projector) off() {
	fmt.Println("Projector is off")
}

// 子系统：放大器
type Amplifier struct {
	volume int
}

func (a *Amplifier) on() {
	fmt.Println("Amplifier is on")
}

func (a *Amplifier) setVolume(volume int) {
	a.volume = volume
	fmt.Println("Amplifier volume set to:", volume)
}

func (a *Amplifier) off() {
	fmt.Println("Amplifier is off")
}

// 外观：家庭影院Facade
type HomeTheaterFacade struct {
	dvd       *DVDPlayer
	projector *Projector
	amplifier *Amplifier
}

func NewHomeTheaterFacade(dvd *DVDPlayer, projector *Projector, amplifier *Amplifier) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvd:       dvd,
		projector: projector,
		amplifier: amplifier,
	}
}

func (h *HomeTheaterFacade) watchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")

	// 协调子系统：打开设备并设置
	h.amplifier.on()
	h.amplifier.setVolume(5)

	h.projector.on()
	h.projector.setInput("DVD")

	h.dvd.on()
	h.dvd.play(movie)

	fmt.Println("Enjoy the movie!")
}

func (h *HomeTheaterFacade) endMovie() {
	fmt.Println("Shutting down the theater...")

	h.dvd.off()
	h.projector.off()
	h.amplifier.off()

	fmt.Println("Theater is closed.")
}

// 外观模式是一种结构型模式，外观模式通过为复杂系统提供一个简单的接口，来简化客户端与子系统之间的交互
func TestFacade(t *testing.T) {
	// 创建子系统实例
	dvd := &DVDPlayer{}
	projector := &Projector{}
	amplifier := &Amplifier{}

	// 创建Facade
	theater := NewHomeTheaterFacade(dvd, projector, amplifier)

	// 客户端只需调用简单方法
	theater.watchMovie("Inception")
	fmt.Println("---")
	theater.endMovie()
}
