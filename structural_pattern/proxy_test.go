package main

import (
	"fmt"
	"testing"
)

// 代理模式就是为被代理的对象提供间接的访问方式，从而可以在访问前后增加额外的处理逻辑

// 比如模拟文件访问服务，代理对象在访问文件之前先检查文件权限

type File interface {
	ReadFile(user string) string
}

type RealFile struct {
	fileName string
}

func (r *RealFile) ReadFile(user string) string {
	return fmt.Sprintf("Reading file %s for user %s", r.fileName, user)
}

type FileProxy struct {
	file         RealFile
	allowedUsers map[string]bool
}

func NewFileProxy(filename string) *FileProxy {
	return &FileProxy{
		file: RealFile{fileName: filename},
		allowedUsers: map[string]bool{
			"admin": true,
			"user":  true,
		},
	}
}

func (f *FileProxy) ReadFile(user string) string {
	if f.allowedUsers[user] {
		return f.file.ReadFile(user)
	}
	return fmt.Sprintf("User %s is not allowed to read file %s", user, f.file.fileName)

}

func TestProxy(t *testing.T) {
	// 创建代理对象
	proxy := NewFileProxy("test.txt")
	fmt.Println(proxy.ReadFile("admin"))
	fmt.Println(proxy.ReadFile("user"))
	fmt.Println(proxy.ReadFile("guest"))
}
