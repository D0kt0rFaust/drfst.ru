/*
	Элемент заказа
*/

package models

import (
	"log"

	"gorm.io/gorm"
)

// Структура сущности
type OrderItem struct {
	gorm.Model
	Title string
}

// Переменные

var DB *gorm.DB

var orderItem OrderItem
var orderItems []OrderItem

// Методы

// Автоматическая миграция из структуры
func OrderItemAutoMigrate() {
	log.Println("OrderItemAutoMigrate")
	DB.AutoMigrate(&OrderItem{})
}

// Добавление записи
func OrderItemCreate(title string) {
	log.Println("OrderItemCreate")
	entity := OrderItem{Title: title}
	DB.Create(&entity)
}

// Список всех записей
func OrderItemList() []OrderItem {
	log.Println("OrderItemList")
	DB.Find(&orderItems)
	return orderItems
}