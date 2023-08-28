package mftb

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startMSG = "/start"
)

func IsStartMsg(message *tgbotapi.Message) bool {
	return IsCommand(message) && strings.Contains(message.Text, startMSG)
}

func IsCommand(message *tgbotapi.Message) (isGood bool) {
	if IsTextmsg(message) {
		isGood = strings.Contains(message.Text, "/")
	}

	return isGood
}

func IsTextmsg(message *tgbotapi.Message) bool {
	return message.Text != ""
}

func IsDocumentMsg(message *tgbotapi.Message) bool {
	return message.Document != nil
}

func CheckUser(id int64) bool {
	return false
}
