package usecase

import (
	"github.com/LittleLe6owski/FilmRecom/internal/buffer"
	"github.com/LittleLe6owski/FilmRecom/internal/model"
	"log"
	"sync"
)

type recFunc func(int64, int64)
type recStruct struct {
	fIlmId int64
	chatId int64
}

type Recommendation struct {
	movieUseCase  model.MovieUseCase
	staffUseCase  model.PersonUseCase
	buffer        buffer.Buffer
	recommendFunc map[string]recFunc
}

const (
	basedOnStaff = "basedOnStaff"
	basedOnGenre = "basedOnGenre"

	NumberOfThreads = 4
)

func NewRecommendationMachine(movieUseCase model.MovieUseCase, staffUseCase model.PersonUseCase,
	buffer buffer.Buffer) model.Recommendation {
	rec := Recommendation{
		movieUseCase: movieUseCase,
		staffUseCase: staffUseCase,
		buffer:       buffer,
	}
	rec.recommendFunc = map[string]recFunc{
		basedOnGenre: rec.recommendBasedOnGenre,
		basedOnStaff: rec.recommendBasedOnStaff,
	}
	return &rec
}

//StartRecommendingFilms Запускает процесс наполнение буффера рекомендованными фильмами.
func (r *Recommendation) StartRecommendingFilms(chatId int64, movieIds ...int64) {

	numThreads := r.calculateNumberThreads(len(movieIds))
	var wg sync.WaitGroup
	jobs := make(chan recStruct)

	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go r.recomWorker(&wg, jobs)
	}

	for _, id := range movieIds {
		jobs <- recStruct{
			fIlmId: id,
			chatId: chatId,
		}
	}
	wg.Wait()
}

//Получает создателей фильма, затем список фильмов, в создании которых он(а) принимал(а) участие, и пишет в буфер.
func (r *Recommendation) recommendBasedOnStaff(fIlmId, chatId int64) {
	people, err := r.staffUseCase.GetStaffByFilmId(fIlmId)
	if err != nil {
		log.Println("")
		return
	}
	for i := range people {
		if people[i].Profession == model.Actor || people[i].Profession == model.Director {
			person, err := r.staffUseCase.GetPersonById(people[i].Id)
			if err != nil {
				log.Println("")
				return
			}
			r.buffer.AppendMovies(chatId, person.Movies)
		}
	}
}

func (r *Recommendation) recommendBasedOnGenre(fIlmId, chatId int64) {
	movies, err := r.movieUseCase.GetSimilarById(fIlmId)
	if err != nil {
		log.Println("")
		return
	}
	r.buffer.AppendMovies(chatId, movies)
}

func (r *Recommendation) calculateNumberThreads(numberFilms int) int {
	if numberFilms > NumberOfThreads {
		return NumberOfThreads
	}
	return numberFilms
}

func (r *Recommendation) recomWorker(wg *sync.WaitGroup, jobs <-chan recStruct) {
	for j := range jobs {
		for i := range r.recommendFunc {
			r.recommendFunc[i](j.fIlmId, j.chatId)
		}
		wg.Done()
	}
}
