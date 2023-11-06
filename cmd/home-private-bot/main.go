package main

import (
	// "context"
	"log"
	"os"
	// "strconv"
	"time"

	"drfst.ru/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var botApiKey string
var DB *gorm.DB

// Инициализируем приложение
func init() {
	// Загружаем env-переменные
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botApiKey = os.Getenv("BOT_API_KEY")

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
}

func main() {
	// Создаём бота с ключом из botApiKey
	bot, err := tgbotapi.NewBotAPI(botApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	        t := time.Now()

			// Создаём элемент заказа
			models.ProductCreate(update.Message.Text)

	        welcomeMessage := "Привет, " +
	            update.Message.Chat.FirstName +  " " +
	            update.Message.Chat.LastName + "!\r\n\r\n" +
				"Элемент заказа '" + update.Message.Text + "' успешно создан\r\n\r\n" +
	            "Техническая информация:\r\n" +
	            "Сообщение: " + update.Message.Text + "\r\n" +
	            "Время: " + t.Format("2006-01-02 15:04:05 UTC")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMessage)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
