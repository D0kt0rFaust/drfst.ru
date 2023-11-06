/*
	Элемент заказа
*/

package models

import (
	"log"

	"gorm.io/gorm"
)

// Структура сущности
type Product struct {
	gorm.Model
	Title string
}

// Переменные

var DB *gorm.DB

var product Product
var products []Product

// Методы

// Автоматическая миграция из структуры
func ProductAutoMigrate() {
	log.Println("ProductAutoMigrate")
	DB.AutoMigrate(&Product{})
}

// Добавление записи
func ProductCreate(title string) {
	log.Println("ProductCreate")
	entity := Product{Title: title}
	DB.Create(&entity)
}

// Список всех записей
func ProductList() []Product {
	log.Println("ProductList")
	DB.Find(&products)
	return products
}