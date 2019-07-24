package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var WorkAnswers = [...]string{"Ты хотел сказать \"чиллить\"?", "Ну я лично скоро на Охотном чилить буду", "Опять на работу пиздос"}

func RandomizeAnswers(Collection [3]string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	AnsIndex := rand.New(s1).Intn(len(Collection))
	return Collection[AnsIndex]
}
func main() {
	proxyUrl, err := url.Parse("socks5://127.0.0.1:9050") // Proxy pass
	if err != nil {
		log.Panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}
	// Auth section
	bot, err := tgbotapi.NewBotAPIWithClient("token", client) //TODO: configs->config.yaml
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account: %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "about":
				msg.Text = "Меня зовут Гореси. Если вы считаете, что я имею отношение к любому реальному человеку, или являюсь какой-либо отсылкой, то вы, наверное, долбоёб"
			case "fact":
				msg.Text = "Ну вообще-то субъективно объективно фактически это моё мнение."
			default:
				msg.Text = "Error: ты пидор, нет такой команды"
			}
			bot.Send(msg)
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		switch {
		case strings.Contains(msg.Text, "работать"):
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = RandomizeAnswers(WorkAnswers)
			bot.Send(msg)
		case strings.Contains(msg.Text, "финик"):
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "Ну вообще-то финик заебал меня хейтить"
			bot.Send(msg)
		case strings.Contains(msg.Text, "залупин"):
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "Залупин бы сначала научился микротик настраивать"
			bot.Send(msg)
		case strings.Contains(msg.Text, "справедливо"):
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "имхо нет"
			bot.Send(msg)
		case strings.Contains(msg.Text, "хорошо плохо"):
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "Ну они с фиником вообще-то меня тупо хейтят постоянно"
			bot.Send(msg)
		}
	}
}
