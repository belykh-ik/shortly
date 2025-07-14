package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"tgBot/broker"
	"tgBot/consts"
	"tgBot/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	servers := []string{"kafka0:9092", "kafka1:9092"}

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

	// Создаем канал для синхронизации консюмера и бота
	in := make(chan string, 1)

	// Создаем консюмера
	go func() {
		consumeer, err := broker.NewConsumer(servers, consts.GroupId, consts.ConsumerTopic)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		log.Println("Start Consumer...")
		consumeer.Start(in)
	}()

	// Ждем старт консюмера
	time.Sleep(3 * time.Second)

	for update := range updates {
		if update.Message == nil { // пропускаем нечитаемые апдейты
			continue
		}
		text := update.Message.Text
		// Ищем все ссылки в тексте сообщения
		links := urlRegex.FindAllString(text, -1)
		if len(links) > 0 {
			for _, link := range links {
				// Ждем ответ максимум 2 секунды
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))
				err := broker.Produce(link)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка отправки")
					bot.Send(msg)
					// Отмена контекста
					cancel()
				} else {
					// Ответ пользователю
					select {
					case data := <-in:
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, data)
						bot.Send(msg)
						// Отмена контекста
						cancel()
					case <-ctx.Done():
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка, попробуйте снова...")
						bot.Send(msg)
					}
				}
			}
		} else {
			// Если ссылок нет
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, отправьте ссылку.")
			bot.Send(msg)
		}
	}
}
