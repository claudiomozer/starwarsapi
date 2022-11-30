package dataprotocols

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type LoadFilmsByUrlRepository interface {
	Load(url string) ([]domaindto.FilmDTO, error)
}
