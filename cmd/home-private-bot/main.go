package main

import (
	// "context"
	"log"
	"os"
	// "strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var botApiKey string

// Инициализируем приложение
func init() {
	// Загружаем env-переменные
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botApiKey = os.Getenv("BOT_API_KEY")
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

	        welcomeMessage := "Привет, " +
	            update.Message.Chat.FirstName +  " " +
	            update.Message.Chat.LastName + "!\r\n\r\n" +
	            "Техническая информация:\r\n" +
	            "Сообщение: " + update.Message.Text + "\r\n" +
	            "Время: " + t.Format("2006-01-02 15:04:05 UTC")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMessage)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
