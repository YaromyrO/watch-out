package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	defaultTimeout = 60
)

type Bot struct {
	api     *tgbotapi.BotAPI
	timeout int
}

func NewBot(apiToken string, timeout int) *Bot {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panic(err)
	}

	if timeout == 0 {
		timeout = defaultTimeout
	}

	return &Bot{
		api:     bot,
		timeout: timeout,
	}
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ПО НАЗВІ"),
		tgbotapi.NewKeyboardButton("ПО КЛЮЧОВОМУ СЛОВУ"),
	),
)

func (bot *Bot) HandleRequests() {
	uc := newUpdateConfig(bot.timeout)
	updatesChan := bot.api.GetUpdatesChan(uc)
	bot.processUpdates(updatesChan)
}

func newUpdateConfig(timeout int) tgbotapi.UpdateConfig {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = timeout

	return uc
}

func (bot *Bot) processUpdates(updatesChan tgbotapi.UpdatesChannel) {
	for update := range updatesChan {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.Command() == "start" {
			msg.Text = startMessage
			msg.ReplyMarkup = numericKeyboard
		} else {
			msg.Text = "I don't know that command"
		}

		if _, err := bot.api.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
