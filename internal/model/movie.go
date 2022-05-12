package model

import tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Movie struct {
	Name        string `json:"nameRu"`
	Id          int64  `json:"filmId"`
	Description string `json:"description"`
	//Countries       []string `json:"countries"`
	//Genres          []string `json:"genres"`
	Rating          string `json:"rating"`
	RatingVoteCount int64  `json:"ratingVoteCount"`
	PosterUrl       string `json:"posterUrl"`
	Staff           Person `json:"-"`
}

type ResponseByKeyword struct {
	Keyword   string  `json:"keyword"`
	PageCount string  `json:"pageCount"`
	Movies    []Movie `json:"films"`
}

type MovieHandler interface {
	HandleButton(int64, string)
	HandleMessage(int64, string)
}

type MovieUseCase interface {
	GetByKeyword(string) ([]Movie, error)
	GetById(int64) (Movie, error)
	GetSimilarById(int64) ([]Movie, error)
	//GetSimilarByFilter(map[string]string) ([]Movie, error)
	//GetPosterById(int64) (string, error)
	//GetTeaserById(int64) (string, error)
}

type MovieLogic interface {
	PrimaryHandle(string, int64) (tgBotApi.Chattable, error)
	ButtonHandle(string, int64) (tgBotApi.Chattable, error)
}
