package main

// 建造者模式就是用来创建复杂对象的一种模式，以一种类似于流式的方式创建对象，
// 每次设置一个属性，然后返回建造者对象本身，这样对象创建的过程就很清晰，最后通过Build函数创建对象。
import (
	"fmt"
	"testing"
	"time"
)

type HttpReqConfig struct {
	URL     string
	Method  string
	Headers map[string]string
	Timeout time.Duration
}

func (c *HttpReqConfig) showConfig() string {
	return fmt.Sprintf("url:%s,method:%s,headers:%v,timeout:%d", c.URL, c.Method, c.Headers, c.Timeout)
}

type HttpReqBuilder struct {
	config *HttpReqConfig
}

func NewHttpReqBuilder(url string) *HttpReqBuilder {
	return &HttpReqBuilder{
		config: &HttpReqConfig{
			URL:     url,
			Method:  "GET",
			Headers: make(map[string]string),
			Timeout: 10 * time.Second,
		},
	}
}

func (b *HttpReqBuilder) SetMethod(method string) *HttpReqBuilder {
	b.config.Method = method
	return b
}
func (b *HttpReqBuilder) AddHeader(key, value string) *HttpReqBuilder {
	b.config.Headers[key] = value
	return b
}

func (b *HttpReqBuilder) SetTimeout(timeout time.Duration) *HttpReqBuilder {
	b.config.Timeout = timeout
	return b
}

func (b *HttpReqBuilder) Build() *HttpReqConfig {
	return b.config
}

func TestBuilder(t *testing.T) {
	simpleConfig := NewHttpReqBuilder("https://www.baidu.com").
		SetMethod("POST").
		AddHeader("Content-Type", "application/json").
		SetTimeout(5 * time.Second).Build()

	fmt.Println(simpleConfig.showConfig())
}
