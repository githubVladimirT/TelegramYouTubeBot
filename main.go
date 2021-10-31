package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"telegram.api/telegram_youtube/youtube"
)

func main() {
	botToken := "2096039520:AAGXsByirzENLmK1i-7X3CQnvfwCzntgJno" // Token of telegram bot@githubVladimirT
	botApi := "https://api.telegram.org/bot" // Bot api@githubVladimirT
	botUrl := botApi + botToken // Bot Url = botApi + botToken@githubVladimirT
	offset := 0

	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Panic(err)
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
		log.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse

	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

func respond(botUrl string, update Update) (error) {
	var botMessage BotMessage

	botMessage.ChatId = update.Message.Chat.ChatId

	videoUrl, err := youtube.GetLastVideo(update.Message.Text)
	if err != nil {
		return err
	}

	botMessage.Text = videoUrl
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}
