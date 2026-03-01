这是一个使用go语言实现22种设计模式的相关代码库[设计模式无定式]
---
### 创建型模式
1. 单例模式
   单例模式指的是全局只存在一个实例，可以用于配置中心、日志对象、全局唯一的线程池或者数据库连接池，全局计数器等。go语言很方便的可以使用sync.Once实现单例（once.Do()方法）
2. 工厂方法模式和抽象工厂模式
   工厂方法模式更关注单个对象及其变体的创建，而抽象工厂模式则更加关注产品族的概念，就是说一个大对象包含多个不同的小对象，而不同的大对象之间又有所不同。
3. 
### 结构型模式

### 行为型模式
1. 命令模式
   命令模式是将一个对象及其方法进行进一步的封装。

### Options模式
   Options模式在Go语言中可以很方便的实现新配置覆盖默认配置，很方便很实用.
   
```go
type Server struct {
   Addr string
   Port int
}

type Option func(*Server)

func WithAddr(addr string) Option {
   return func(s *Server) {
      s.Addr = addr
   }
}
func WithPort(port string) Option {
   return func(s *Server) {
      s.Port = port
   }
}

func NewServer(opts ...Options) *Server {
   s := &Server {
      s.Addr: "127.0.0.1",
      s.Port: 8888,
   }
   for _, opt := range opt {
      opt(s)
   }

   return s
}

// s := NewServer(WithPort("6666"))
```
