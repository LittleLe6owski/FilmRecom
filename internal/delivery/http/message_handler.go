package http

import (
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

type chatConfig func(int64)

type MessageHandler struct {
	messageCh      tgBotApi.UpdatesChannel
	api            *tgBotApi.BotAPI
	movieHandler   model.MovieHandler
	messageBuilder model.MessageBuilder
	handlers       map[string]chatConfig
}

func NewMessageHandler(messageCh tgBotApi.UpdatesChannel, api *tgBotApi.BotAPI,
	movieHandler model.MovieHandler, messageBuilder model.MessageBuilder) model.MessageHandler {

	mh := MessageHandler{
		messageCh:      messageCh,
		api:            api,
		movieHandler:   movieHandler,
		messageBuilder: messageBuilder,
	}
	mh.handlers = map[string]chatConfig{
		model.Start: mh.printStartMessage,
		model.Help:  mh.printHelp,
	}
	return &mh
}

func (m *MessageHandler) RunHandleMessages() {
	for update := range m.messageCh {
		if update.CallbackQuery != nil {
			m.movieHandler.HandleButton(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			continue
		}
		if update.Message != nil {
			switch _, isExist := m.handlers[update.Message.Text]; isExist {
			case false:
				m.movieHandler.HandleMessage(update.Message.Chat.ID, update.Message.Text)
			default:
				m.handlers[update.Message.Text](update.Message.Chat.ID)
			}
		}
	}
}

func (m *MessageHandler) printStartMessage(chatId int64) {
	msg := m.messageBuilder.BuildMessage(model.Message{
		Text: viper.GetString(model.StartMessage),
	}, chatId, false)

	_, err := m.api.Send(msg)
	if err != nil {
		log.Printf(notSent, chatId)
	}
}

func (m *MessageHandler) printHelp(chatId int64) {
	msg := m.messageBuilder.BuildMessage(model.Message{
		Text: viper.GetString(model.HelpMessage),
	}, chatId, false)

	_, err := m.api.Send(msg)
	if err != nil {
		log.Printf(notSent, chatId)
	}
}
