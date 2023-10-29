/*
	Список типовых продуктов для заказа.
	Будет использоваться для унификации заявок на покупку.
*/

package models

import (
	"log"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
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
func ProductCreate() {
	log.Println("ProductCreate")
	entity := Product{Title: "Test"}
	DB.Create(&entity)
}

// Список всех продуктов
func ProductList() {
	log.Println("ProductList")
	result := DB.Find(&products)
	log.Println(result.RowsAffected)
}