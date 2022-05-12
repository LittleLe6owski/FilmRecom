package internal

import (
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"github.com/LittleLe6owski/FilmRecom/pkg/telegram"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type MessageBuilder struct {
	buttonBuilder telegram.ButtonBuilder
}

func NewMessageBuilder(buttonBuilder telegram.ButtonBuilder) model.MessageBuilder {
	return &MessageBuilder{
		buttonBuilder: buttonBuilder,
	}
}

func (m *MessageBuilder) BuildMessage(message model.Message, chatId int64,
	fully bool, buttons ...map[string]string) tgBotApi.Chattable {
	if message.Text == "" {
		if fully {
			message.Text = m.genLongText(message.Movie)
		}
		message.Text = m.genShortText(message.Movie)
	}
	msg := tgBotApi.NewMessage(chatId, message.Text)
	if len(buttons) != 0 {
		msg.ReplyMarkup = m.genButtons(buttons[0], chatId)
	}
	return msg
}

func (m *MessageBuilder) genShortText(movies []model.Movie) (messageText string) {
	for i := range movies {
		messageText += model.NameMessage
		messageText += movies[i].Name
		messageText += model.Description
		messageText += movies[i].Description
	}
	return
}

func (m *MessageBuilder) genLongText(movies []model.Movie) (messageText string) {
	messageText = m.genShortText(movies)
	for i := range movies {
		messageText += model.Rating
		messageText += movies[i].Rating
		messageText += model.RatingVoteCount
		messageText += strconv.FormatInt(movies[i].RatingVoteCount, 10)
	}
	return
}

func (m *MessageBuilder) genButtons(buttons map[string]string, chatId int64) tgBotApi.InlineKeyboardMarkup {
	repMark := make([][]telegram.Button, 1)
	for i, key := range buttons {
		repMark[0] = []telegram.Button{{
			Text:     key,
			Callback: strconv.FormatInt(chatId, 10) + "_" + buttons[i],
		}}
	}
	return m.buttonBuilder.BuildKeyboard(repMark)
}
