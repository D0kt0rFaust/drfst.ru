/*
	Список типовых продуктов для заказа.
	Будет использоваться для унификации заявок на покупку.
*/

package models

import (
	"log"

	"gorm.io/gorm"
)

// Структура сущности
type ProductTemplate struct {
	gorm.Model
	Title string
}

// Переменные

// var DB *gorm.DB

var productTemplate ProductTemplate
var productTemplates []ProductTemplate

// Методы

// Автоматическая миграция из структуры
func ProductTemplateAutoMigrate() {
	log.Println("ProductTemplateAutoMigrate")
	DB.AutoMigrate(&ProductTemplate{})
}

// Добавление записи
func ProductTemplateCreate() {
	log.Println("ProductTemplateCreate")
	entity := ProductTemplate{Title: "Test"}
	DB.Create(&entity)
}

// Список всех записей
func ProductTemplateList() []ProductTemplate {
	log.Println("ProductTemplateList")
	DB.Find(&productTemplates)
	return productTemplates
}