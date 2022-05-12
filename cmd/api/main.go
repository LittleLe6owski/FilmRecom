package main

import (
	"github.com/LittleLe6owski/FilmRecom/internal"
	"github.com/LittleLe6owski/FilmRecom/internal/buffer"
	"github.com/LittleLe6owski/FilmRecom/internal/delivery/http"
	"github.com/LittleLe6owski/FilmRecom/internal/logic"
	"github.com/LittleLe6owski/FilmRecom/internal/usecase"
	"github.com/LittleLe6owski/FilmRecom/pkg/queryBuilder"
	"github.com/LittleLe6owski/FilmRecom/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

func InitConnect() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("botToken"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.GetUpdatesChan(u), bot
}

func main() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	ch, api := InitConnect()
	buf := buffer.NewBuffer()

	var (
		qb = queryBuilder.NewQueryBuilder(
			viper.GetString("schema"),
			viper.GetString("host"),
			viper.GetStringMapString("requestHeaders"))
		bb = telegram.NewButtonBuilder()
		mb = internal.NewMessageBuilder(bb)
	)

	var (
		movieUse = usecase.NewMovieUseCase(qb, buf)
		staffUse = usecase.NewStaffUseCase(qb)
		recUse   = usecase.NewRecommendationMachine(movieUse, staffUse, buf)
	)

	var (
		movieLogic = logic.NewMovieLogic(movieUse, recUse)
	)

	movieHandler := http.NewMovieHandler(api, movieLogic, mb)
	h := http.NewMessageHandler(ch, api, movieHandler, mb)

	h.RunHandleMessages()
	select {}
}
