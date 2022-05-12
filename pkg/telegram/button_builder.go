package telegram

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
	Конструктор для кнопок бота в телеграм.
*/

type ButtonConstructor struct{}

func NewButtonBuilder() ButtonBuilder {
	return &ButtonConstructor{}
}

func (b ButtonConstructor) BuildKeyboard(buttons [][]Button) (markup tgBotApi.InlineKeyboardMarkup) {

	columnButtons := make([][]tgBotApi.InlineKeyboardButton, 0)
	var lineButtons []tgBotApi.InlineKeyboardButton
	for _, line := range buttons {
		lineButtons = make([]tgBotApi.InlineKeyboardButton, 0)
		for j := range line {
			lineButtons = append(lineButtons, tgBotApi.NewInlineKeyboardButtonData(line[j].Text, line[j].Callback))
		}
		columnButtons = append(columnButtons, lineButtons)
	}
	return tgBotApi.NewInlineKeyboardMarkup(columnButtons...)
}
