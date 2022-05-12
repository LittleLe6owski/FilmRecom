package model

type Recommendation interface {
	StartRecommendingFilms(chatId int64, movieIds ...int64)
}
