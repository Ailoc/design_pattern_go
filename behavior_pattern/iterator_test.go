package main

import (
	"errors"
	"testing"
)

type User struct { // 集合元素
	name string
	age  int
}

type Iterator interface { // 迭代器接口
	HasNext() bool
	Next() (interface{}, error)
}

type Collection interface { // 集合接口，其中包含
	CreateIterator() Iterator // 创建迭代器的方法
}

type UserCollection struct { // 具体的集合类
	Users []User
}

type UserIterator struct { // 具体的迭代器类
	index      int
	collection *UserCollection
}

func (ui *UserIterator) HasNext() bool {
	return ui.index < len(ui.collection.Users)
}

func (ui *UserIterator) Next() (interface{}, error) {
	if ui.HasNext() {
		user := ui.collection.Users[ui.index]
		ui.index++
		return user, nil
	}
	return nil, errors.New("No more elements")
}

func (uc *UserCollection) CreateIterator() Iterator {
	return &UserIterator{index: 0, collection: uc}
}

func TestIterator(t *testing.T) {
	uc := &UserCollection{
		Users: []User{
			{"Mike", 20},
			{"Joe", 30},
			{"Jane", 40},
		},
	}
	it := uc.CreateIterator()
	for it.HasNext() {
		user, _ := it.Next()
		t.Log(user)
	}
	t.Log("Finished iterating")

}
