package main

import (
	"log"
	"os"
	"regexp"
	"tgBot/broker"
	"tgBot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env")
	}
	token := &models.Bot{
		Token: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	bot, err := tgbotapi.NewBotAPI(token.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем получение обновлений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	// Регулярное выражение для обнаружения ссылок
	urlRegex := regexp.MustCompile(`https?://[\w.-]+(?:/[\w./?%&=-]*)?`)

	for update := range updates {
		if update.Message == nil { // пропускаем нечитаемые апдейты
			continue
		}
		text := update.Message.Text
		// Ищем все ссылки в тексте сообщения
		links := urlRegex.FindAllString(text, -1)
		if len(links) > 0 {
			for _, link := range links {
				err := broker.Produce(link)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка, попробуйте снова...")
					bot.Send(msg)
				} else {
					log.Printf("Received link: %s from user %d", link, update.Message.From.ID)
					// Ответ пользователю
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ссылка получена и будет обработана.")
					bot.Send(msg)
				}
			}
		} else {
			// Если ссылок нет
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, отправьте ссылку.")
			bot.Send(msg)
		}
	}
}
