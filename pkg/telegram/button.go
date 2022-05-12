package telegram

import tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	Yes       = "Да"
	No        = "Нет"
	Rollback  = "Вернуться назад"
	Recommend = "Порекомендовать фильм"
	More      = "Посоветовать еще!"
)

type Button struct {
	Text     string
	Callback string
}

type ButtonBuilder interface {
	BuildKeyboard([][]Button) tgBotApi.InlineKeyboardMarkup
}
