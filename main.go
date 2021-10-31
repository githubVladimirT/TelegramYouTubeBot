package main

import (
	"log"

	"telegram.api/telegram_youtube/youtube"
	tel "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tel.NewBotAPI("2096039520:AAGXsByirzENLmK1i-7X3CQnvfwCzntgJno")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tel.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// msg := tel.NewMessage(update.Message.Chat.ID, update.Message.Text)

		videoUrl, err := youtube.GetLastVideo(update.Message.Text, 1)
		if err != nil {
			log.Println(err)
		}
		bot.Send(tel.NewMessage(update.Message.Chat.ID, err.Error()))
		bot.Send(tel.NewMessage(update.Message.Chat.ID, videoUrl))
	}
}