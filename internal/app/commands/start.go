package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Start(inputMsg *tgbotapi.Message) {
	c.client.SendReplyMarkup(inputMsg.Chat.ID, "Сервис запущен ✅",
		c.keyboardService.Keyboard)
}
