package httprepo

import (
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

	return nil, nil
}

func (repository *FilmRepository) isUrlInvalid(url string) bool {
	regex := regexp.MustCompile(`https://swapi.dev/api/films/\d+$`)
	return !regex.MatchString(url)
}
