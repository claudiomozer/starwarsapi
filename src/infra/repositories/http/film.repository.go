package httprepo

import (
	"encoding/json"
	"errors"
	"regexp"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type FilmRepository struct{}

func NewFilmRepository() *FilmRepository {
	return &FilmRepository{}
}

func (repository *FilmRepository) Load(url string) (*domaindto.FilmDTO, error) {

	if repository.isUrlInvalid(url) {
		return nil, errors.New("Impossível buscar filme na API: URL inválida")
	}

	status, body, err := Get(url)
	var filmDTO *domaindto.FilmDTO = &domaindto.FilmDTO{}
	err = json.Unmarshal(body, filmDTO)

	if status == 404 {
		return nil, nil
	}

	return filmDTO, err
}

func (repository *FilmRepository) isUrlInvalid(url string) bool {
	regex := regexp.MustCompile(`https://swapi.dev/api/films/\d+/{0,1}$`)
	return !regex.MatchString(url)
}
