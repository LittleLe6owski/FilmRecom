package buffer

import (
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"sync"
)

type Buffer interface {
	AppendMovies(chatId int64, movies []model.Movie)
	GetMovies(chatId int64) (movie []model.Movie, found bool)
	ClearMovies()
}

//buffer Хранит фильмы, и индекс последнего фильма, который был предложен.
type buffer struct {
	mx        sync.RWMutex
	movies    map[int64][]model.Movie
	lastIndex map[int64]int32
}

func NewBuffer() Buffer {
	return &buffer{
		movies:    make(map[int64][]model.Movie),
		lastIndex: make(map[int64]int32),
	}
}

func (c *buffer) AppendMovies(chatId int64, movies []model.Movie) {
	c.mx.Lock()
	c.mx.Unlock()
	for i := range movies {
		c.movies[chatId] = append(c.movies[chatId], movies[i])
	}
}

func (c *buffer) GetMovies(chatId int64) (movies []model.Movie, found bool) {
	c.mx.Lock()
	defer c.mx.RUnlock()
	movies, found = c.movies[chatId]
	if found {
		//Отдаём 5 элементов, начиная с последнего неизвлечённого(того что мы ещё не отдали пользователю) фильма.
		movies = c.movies[chatId][c.lastIndex[chatId] : c.lastIndex[chatId]+5]
		c.lastIndex[chatId] += 5
	}
	return
}

func (c *buffer) ClearMovies() {
	//TODO implement me
	panic("implement me")
}
