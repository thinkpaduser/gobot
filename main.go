package main

import (
	"flag"
	"github.com/thinkpaduser/gobot/pkg/config"
	"github.com/thinkpaduser/gobot/pkg/static"
	"github.com/thinkpaduser/gobot/pkg/utils"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var conf *config.Config
var mess config.MsgDB

// Returns a random element from array
func Random(Collection []string) string {
	return Collection[RandomInt(len(Collection))]
}

func RandomInt(Range int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed).Intn(Range)
}

// Load configs and YAML msg library
func init() {
	const (
		defaultConfigFilename   = "./configs/conf.yaml"
		defaultMessagesFilename = "./configs/messages.yaml"
	)
	var err error
	var configFilename string
	flag.StringVar(&configFilename, "config", defaultConfigFilename, "the config filename") // Optional app [configfile] launch option for custom configuration
	flag.Parse()
	if err = initConf(configFilename); err != nil {
		log.Panic("Can't init config: %s", err.Error())
	}
	if err = initMsgDB(defaultMessagesFilename); err != nil {
		log.Panic("Can't init messages storage: %s", err.Error())
	}
}

func initConf(filename string) (err error) { // Importing config structure
	if conf, err = config.NewConf(filename); err != nil {
		return
	}
	return
}

func initMsgDB(filename string) (err error) { // Importing messages hash table
	if mess, err = config.NewMsgDB(filename); err != nil {
		return
	}
	return
}

// Get an answer from msg library by keys
func FindAnswer(m map[string][]string, s string, usr string) string {
	var res string
	var keys []string

	for k := range m {

		if strings.Contains(strings.ToLower(s), k) {
			keys = append(keys, k)
		}

		isQuestion, _ := regexp.MatchString(`.*\?$`, s)
		if isQuestion && usr == conf.Conf.Target {
			res = "дай аналог лол"
		}

		if keys != nil {
			key := Random(keys)
			res = Random(m[key])
		}

	}
	return res
}

// Get a chance of answering or not
func GetChance(k int, s string) string {
	var res string
	if RandomInt(10)*k/10 >= 5 {
		res = s
	}
	return res
}

// Append to a slice distinct values
func AppendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func main() {
	var members []string

	client := utils.NewTorClient()
	bot, err := tgbotapi.NewBotAPIWithClient(conf.Conf.Token, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Fatalf("can't init a client: %s", err.Error())
	}
	//bot.Debug = true
	log.Printf("Authorized on account: %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("can't get updates: %s", err.Error())
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		lastMsgTime := time.Now()
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "about":
				msg.Text = static.Commands.ABOUT
			case "abuse":
				msg.Text = static.Commands.ABUSE
			default:
				msg.Text = ""
			}
			bot.Send(msg)
		}

		usrID := update.Message.From.String()
		members = AppendIfMissing(members, usrID)
		log.Println(members)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		go func(lastMsg time.Time) {
			for t := range time.NewTicker(time.Minute * 60).C {
				if time.Since(lastMsg) > 80 * time.Minute {
					continue
				}
				luckyMan := Random(members)
				msg.Text = "@" + luckyMan + " чекни ЛС"
				log.Printf("Ticker, %v", t)
				bot.Send(msg)
			}
		}(lastMsgTime)

		tempMultiplier := 9 // TODO: refactor or remove this useless crap
		if GetChance(tempMultiplier, FindAnswer(mess, msg.Text, usrID)) != "" {
			msg.Text = FindAnswer(mess, msg.Text, usrID)
			bot.Send(msg)
		}
	}
}
