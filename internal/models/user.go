/*
	Пользователь системы
*/

package models

import (
	"log"

	"gorm.io/gorm"
)

// Структура сущности
type User struct {
	gorm.Model
	Firstname string
	Lastname string
	Middlename string
	TelegramId int
	TelegramUsername string
	Approved bool
	Description string
}

// Переменные
var user User
var users []User

// Методы

// Автоматическая миграция из структуры
func UserAutoMigrate() {
	log.Println("UserAutoMigrate")
	DB.AutoMigrate(&User{})
}

// Добавление записи
func UserCreate(description string) {
	log.Println("UserCreate")
	entity := User{Description: description}
	DB.Create(&entity)
}

// Список всех записей
func UserList() []User {
	log.Println("UserList")
	DB.Find(&users)
	return users
}