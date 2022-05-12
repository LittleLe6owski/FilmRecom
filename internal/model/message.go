package model

import tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message struct {
	Text  string
	Image string
	Movie []Movie
	Staff Person
}

const (
	Start = "/start"
	Help  = "/help"

	HelpMessage  = "helpMessage"
	StartMessage = "startMessage"

	NameMessage     = "Название фильма:"
	Description     = "Краткое описание:"
	Rating          = "Рейтинг фильма:"
	RatingVoteCount = "Количество оценок:"
)

type MessageHandler interface {
	RunHandleMessages()
}

type MessageBuilder interface {
	BuildMessage(Message, int64, bool, ...map[string]string) tgBotApi.Chattable
}
