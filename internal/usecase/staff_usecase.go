package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"github.com/LittleLe6owski/FilmRecom/pkg/queryBuilder"
	"strconv"
)

const (
	getStaffByFilmId   = "staff"
	getStaffByPersonId = "films/%d"
)

type StaffUseCase struct {
	queryBuilder queryBuilder.QueryBuilder
}

func NewStaffUseCase(builder queryBuilder.QueryBuilder) model.PersonUseCase {
	return &StaffUseCase{
		queryBuilder: builder,
	}
}

func (s StaffUseCase) GetStaffByFilmId(id int64) ([]model.Person, error) {
	body, err := s.queryBuilder.Get(queryBuilder.RequestParams{
		Version:  "/api/v1/",
		Path:     getStaffByFilmId,
		RawQuery: map[string]string{"filmId": strconv.FormatInt(id, 10)},
	})
	if err != nil {
		return nil, err
	}
	listFilms := make([]model.Person, 0)
	err = json.Unmarshal(body, &listFilms)
	if err != nil {
		return nil, err
	}
	return listFilms, err
}

func (s StaffUseCase) GetPersonById(id int64) (person model.Person, err error) {
	body, err := s.queryBuilder.Get(queryBuilder.RequestParams{
		Version: "/api/v1/",
		Path:    fmt.Sprintf(getStaffByPersonId, id),
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &person)
	if err != nil {
		return
	}
	return
}
