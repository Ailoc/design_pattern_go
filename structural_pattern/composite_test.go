package main

import (
	"errors"
	"fmt"
	"testing"
)

// 组件接口：FileSystemNode
type FileSystemNode interface {
	Name() string
	Size() int64 // 文件/目录大小（字节）
	Add(node FileSystemNode) error
	Remove(name string) error
	List() []FileSystemNode // 返回子节点列表（目录用，叶子返回nil）
}

// 叶子：文件
type File1 struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File1 {
	return &File1{name: name, size: size}
}

func (f *File1) Name() string {
	return f.name
}

func (f *File1) Size() int64 {
	return f.size
}

func (f *File1) Add(node FileSystemNode) error {
	return errors.New("files cannot have children")
}

func (f *File1) Remove(name string) error {
	return errors.New("files cannot have children")
}

func (f *File1) List() []FileSystemNode {
	return nil
}

// 组合：目录
type Directory struct {
	name     string
	children map[string]FileSystemNode // 用map便于按名查找
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make(map[string]FileSystemNode),
	}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int64 {
	total := int64(0)
	for _, child := range d.children {
		total += child.Size() // 递归计算子节点大小
	}
	return total
}

func (d *Directory) Add(node FileSystemNode) error {
	if _, exists := d.children[node.Name()]; exists {
		return fmt.Errorf("child %s already exists", node.Name())
	}
	d.children[node.Name()] = node
	return nil
}

func (d *Directory) Remove(name string) error {
	if _, exists := d.children[name]; !exists {
		return fmt.Errorf("child %s not found", name)
	}
	delete(d.children, name)
	return nil
}

func (d *Directory) List() []FileSystemNode {
	nodes := make([]FileSystemNode, 0, len(d.children))
	for _, child := range d.children {
		nodes = append(nodes, child)
	}
	return nodes
}

// 辅助函数：递归打印树结构（展示遍历）
func printTree(node FileSystemNode, prefix string) {
	fmt.Printf("%s%s (size: %d bytes)\n", prefix, node.Name(), node.Size())
	if dir, ok := node.(*Directory); ok {
		for _, child := range dir.List() {
			printTree(child, prefix+"  ")
		}
	}
}

func TestComposite(t *testing.T) {
	// 创建叶子文件
	file1 := NewFile("document.txt", 1024)
	file2 := NewFile("image.jpg", 2048)
	file3 := NewFile("script.go", 512)

	// 创建目录
	docsDir := NewDirectory("docs")
	imagesDir := NewDirectory("images")

	// 构建树
	root := NewDirectory("home")
	root.Add(docsDir)
	root.Add(imagesDir)

	docsDir.Add(file1)
	docsDir.Add(file3)
	imagesDir.Add(file2)

	// 打印树结构（递归）
	fmt.Println("File system tree:")
	printTree(root, "")

	// 计算总大小
	fmt.Printf("\nTotal size of home directory: %d bytes\n", root.Size())

	// 移除文件并验证
	if err := docsDir.Remove("document.txt"); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nAfter removing document.txt:")
	printTree(root, "")
}
