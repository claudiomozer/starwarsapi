package dataprotocols

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type CreateFilmRepository interface {
	Create(filmDTO *domaindto.FilmDTO) (id string, err error)
}
