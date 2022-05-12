package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"github.com/LittleLe6owski/FilmRecom/pkg/queryBuilder"
	"strconv"
)

const (
	getByKeywordPath = "films/search-by-keyword/"
	getByIdPath      = "films/"
	getSimilarById   = "films/%d/similars/"
	//getTeaserById    = ""
	//getPosterById    = ""
)

type MovieUseCase struct {
	queryBuilder queryBuilder.QueryBuilder
}

func NewMovieUseCase(builder queryBuilder.QueryBuilder) model.MovieUseCase {
	return &MovieUseCase{
		queryBuilder: builder,
	}
}

//GetByKeyword Получить список model.Movie по ключевому слову
func (m MovieUseCase) GetByKeyword(value string) ([]model.Movie, error) {
	body, err := m.queryBuilder.Get(queryBuilder.RequestParams{
		Version:  "/api/v2.1/",
		Path:     getByKeywordPath,
		RawQuery: map[string]string{"keyword": value},
	})
	if err != nil {
		return nil, err
	}
	listFilms := new(model.ResponseByKeyword)
	err = json.Unmarshal(body, listFilms)
	if err != nil {
		return nil, err
	}
	return listFilms.Movies, nil
}

//GetById Получить полную информацию о фильме по его id
func (m MovieUseCase) GetById(id int64) (model model.Movie, err error) {
	body, err := m.queryBuilder.Get(queryBuilder.RequestParams{
		Version: "/api/v2.2/",
		Path:    getByIdPath + strconv.FormatInt(id, 10),
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return
	}
	return
}

//GetSimilarById Получить фильмы, похожие на тот, id которого передали
func (m MovieUseCase) GetSimilarById(id int64) ([]model.Movie, error) {
	body, err := m.queryBuilder.Get(queryBuilder.RequestParams{
		Version: "/api/v2.1/",
		Path:    fmt.Sprintf(getSimilarById, id),
	})
	if err != nil {
		return nil, err
	}
	listFilms := new(model.ResponseByKeyword)
	err = json.Unmarshal(body, listFilms)
	if err != nil {
		return nil, err
	}
	return listFilms.Movies, nil
}
