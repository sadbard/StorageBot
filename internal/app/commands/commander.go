package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sadbard/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, prSrvc *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: prSrvc,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
			"Data: "+update.CallbackQuery.Data)
		c.bot.Send(msg)
		return
	}

	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}