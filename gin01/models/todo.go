package models

import (
	"github.com/jinzhu/gorm"
	"haimait/learn/gin01/global"
)

// Todo Model
type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	err = global.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = global.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = global.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = global.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = global.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
