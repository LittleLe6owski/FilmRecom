package http

import (
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	notSent      = "User with chatId = %d, not sent message"
	errorMessage = "Processing error, try letter"
)

type MovieHandler struct {
	api            *tgBotApi.BotAPI
	movieLogic     model.MovieLogic
	messageBuilder model.MessageBuilder
}

func NewMovieHandler(api *tgBotApi.BotAPI, movieLogic model.MovieLogic,
	messageBuilder model.MessageBuilder) model.MovieHandler {
	return &MovieHandler{
		api:            api,
		movieLogic:     movieLogic,
		messageBuilder: messageBuilder,
	}
}

func (m *MovieHandler) HandleButton(chatId int64, callBack string) {
	msg, err := m.movieLogic.ButtonHandle(callBack, chatId)
	if err != nil {
		log.Printf("")
		msg = m.messageBuilder.BuildMessage(model.Message{
			Text: errorMessage,
		}, chatId, false)
	}
	_, err = m.api.Send(msg)
	if err != nil {
		log.Printf(notSent, chatId)
	}
}

func (m *MovieHandler) HandleMessage(chatId int64, messageText string) {
	msg, err := m.movieLogic.PrimaryHandle(messageText, chatId)
	if err != nil {
		log.Printf("")
		msg = m.messageBuilder.BuildMessage(model.Message{
			Text: errorMessage,
		}, chatId, false)
	}
	_, err = m.api.Send(msg)
	if err != nil {
		log.Printf(notSent, chatId)
	}
}
