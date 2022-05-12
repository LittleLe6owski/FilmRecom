package model

type Person struct {
	Id            int64   `json:"staffId"`
	Name          string  `json:"nameRu"`
	ProfessionKey string  `json:"professionKey"`
	Profession    string  `json:"profession"`
	PosterUrl     string  `json:"posterUrl"`
	Movies        []Movie `json:"films"`
}

const (
	Director = "Director"
	Actor    = "Actor"
)

type PersonUseCase interface {
	GetStaffByFilmId(int64) ([]Person, error)
	GetPersonById(int64) (Person, error)
}
