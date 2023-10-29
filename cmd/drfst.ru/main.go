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

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello World!</h1>"))
}

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
}

func main() {
    models.ProductAutoMigrate()
	models.ProductCreate()
    models.ProductList()

	port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}