package data

import (
	"errors"
	"fmt"

	dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"
)

type CreateFilmsFromUrlsUseCase struct {
	loadFilmByUrlRepository dataprotocols.LoadFilmByUrlRepository
	createFilmRepository    dataprotocols.CreateFilmRepository
}

func NewCreateFilmsFromUrlsUseCase(
	loadFilmByUrlRepository dataprotocols.LoadFilmByUrlRepository,
	createFilmRepository dataprotocols.CreateFilmRepository,
) *CreateFilmsFromUrlsUseCase {

	return &CreateFilmsFromUrlsUseCase{
		loadFilmByUrlRepository: loadFilmByUrlRepository,
		createFilmRepository:    createFilmRepository,
	}

}

func (usecase *CreateFilmsFromUrlsUseCase) Create(urls []string) (ids []string, err error) {
	for _, url := range urls {
		filmDTO, err := usecase.loadFilmByUrlRepository.Load(url)
		if err != nil {
			return nil, err
		}

		if filmDTO == nil {
			return nil, errors.New(fmt.Sprintf("Erro ao buscar filme: nenhum retorno para a URL: %s", url))
		}

		id, err := usecase.createFilmRepository.Create(filmDTO)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return
}
