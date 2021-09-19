package models

import (
	"bubble/dao"
)

type ToDo struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func CreateTodo(todo *ToDo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func GetAllTodo() (todolist []ToDo, err error) {
	//var todos []dto.ToDoParam
	err = dao.DB.Find(&todolist).Error
	// err = dao.DB.Table("to_dos").Find(&todolist).Error

	// fmt.Printf("todolsit: %v", todolist)
	//for _, val := range todolist {
	//	fmt.Printf("%v", val)
	//}
	return
}

func GetTodo(id string) (todo *ToDo, err error) {
	todo = new(ToDo)
	err = dao.DB.Where("id=?", id).First(todo).Error
	return
}

func UpdateTodo(todo *ToDo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&ToDo{}).Error
	return
}

