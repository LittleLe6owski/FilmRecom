package logic

import (
	"github.com/LittleLe6owski/FilmRecom/internal/buffer"
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"github.com/LittleLe6owski/FilmRecom/pkg/telegram"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type MovieLogic struct {
	movieUse       model.MovieUseCase
	recUse         model.Recommendation
	buffer         buffer.Buffer
	messageBuilder model.MessageBuilder
	buttonBuilder  telegram.ButtonBuilder
	buttonFunc     map[string]func()
}

func NewMovieLogic(movieUse model.MovieUseCase, recUse model.Recommendation) model.MovieLogic {
	ml := MovieLogic{
		movieUse: movieUse,
		recUse:   recUse,
	}
	return &ml
}

func (m MovieLogic) PrimaryHandle(textMessage string, chatId int64) (tgBotApi.Chattable, error) {

	firstSimilar := make([]model.Movie, 0)
	var movieIds string
	for _, name := range strings.Split(textMessage, ",") {
		movies, err := m.movieUse.GetByKeyword(name)
		if err != nil {
			return nil, err
		}
		m.buffer.AppendMovies(-chatId, movies[1:3])
		movieIds += "_" + strconv.FormatInt(movies[0].Id, 10)
		firstSimilar = append(firstSimilar, movies[0])
	}
	return m.messageBuilder.BuildMessage(model.Message{
		Movie: firstSimilar,
	}, chatId, false, map[string]string{
		telegram.Yes: telegram.Yes + movieIds,
		telegram.No:  telegram.No}), nil
}

func (m MovieLogic) ButtonHandle(callBack string, chatId int64) (tgBotApi.Chattable, error) {
	if strings.HasPrefix(callBack, telegram.Yes) {
		movieIds := make([]int64, 0)
		for id := range strings.Split(callBack, "_") {
			movieIds = append(movieIds, int64(id))
		}
		m.recUse.StartRecommendingFilms(chatId, movieIds...)
		time.Sleep(time.Second * 5)
	}
	movie, found := m.buffer.GetMovies(-chatId)
	if !found {
		return m.messageBuilder.BuildMessage(model.Message{
			Text: "Recommend movies not found \n" + viper.GetString(model.HelpMessage),
		}, chatId, false), nil
	}
	return m.messageBuilder.BuildMessage(model.Message{
		Movie: movie,
	}, chatId, true), nil
}
