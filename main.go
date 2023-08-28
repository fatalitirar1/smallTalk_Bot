package main

import (
	dbhendler "ST_bot/DBhendler"
	"ST_bot/mftb"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	token, err := MustGetToken()
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for up := range updates {
		if up.Message != nil {
			if mftb.IsStartMsg(up.Message) {
				SandFirstInstructions(up.Message.Chat.ID, bot)
			}
		}
	}

}

func MustGetToken() (string, error) {
	file, err := os.ReadFile("eternal/token.txt")
	return string(file), err
}

func SandFirstInstructions(id int64, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig
	if !dbhendler.IsChatExists(id) {
		dbhendler.CreateUser(id)
		msg = tgbotapi.NewMessage(id, "hi new one")
	} else {
		msg = tgbotapi.NewMessage(id, "u allready exist")
	}
	bot.Send(msg)

}
