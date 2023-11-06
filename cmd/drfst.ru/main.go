package main

import (
	"log"
	"net/http"
	"os"

    "drfst.ru/internal/models"

    "github.com/joho/godotenv"
    "gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func init() {
    // Загружаем переменные окружения из файла .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Получаем переменные окружения для подключения к основной БД
    mysqlHost := os.Getenv("LOCAL_MYSQL_HOST")
    mysqlPort := os.Getenv("LOCAL_MYSQL_PORT")
    mysqlUsername := os.Getenv("LOCAL_MYSQL_USERNAME")
    mysqlPassword := os.Getenv("LOCAL_MYSQL_PASSWORD")
    mysqlDatabase := os.Getenv("LOCAL_MYSQL_DATABASE")
    
    // Подключаемся
    dsn := mysqlUsername + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Cannot connect mysql")
    }

    // Передаём подключение в модуль models
    models.DB = DB

    // Мигрируем таблички
    models.ProductAutoMigrate()
    models.ProductTemplateAutoMigrate()
    models.UserAutoMigrate()
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/createProduct", createProductHandler)
    mux.HandleFunc("/createProductTemplate", createProductTemplateHandler)

    http.ListenAndServe(":"+port, mux)
}

// Главная страница
func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Выводим список элементов заказа
    productList := models.ProductList()
    w.Write([]byte("### Список элементов заказа\r\n\r\n"))
    for _, value := range productList {
        w.Write([]byte(value.Title + "\r\n"))
    }
    
    // Выводим список продуктов
    productTemplateList := models.ProductTemplateList()
    w.Write([]byte("\r\n### Список типовых продуктов\r\n\r\n"))
    for _, value := range productTemplateList {
        w.Write([]byte(value.Title + "\r\n"))
    }

    // Выводим список пользователей
    userList := models.UserList()
    w.Write([]byte("\r\n### Список пользователей\r\n\r\n"))
    for _, value := range userList {
        w.Write([]byte(value.Lastname + " " + value.Firstname + " " + "\r\n"))
    }
}

// Создание элемента заказа
func createProductHandler(w http.ResponseWriter, r *http.Request) {
    models.ProductCreate("1234")
}

// Создание типового продукта
func createProductTemplateHandler(w http.ResponseWriter, r *http.Request) {
    models.ProductTemplateCreate()
}
