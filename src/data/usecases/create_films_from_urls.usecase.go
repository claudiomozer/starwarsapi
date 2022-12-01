package data

import dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"

type CreateFilmsFromUrlsUseCase struct {
	loadFilmByUrlRepository dataprotocols.LoadFilmByUrlRepository
}

func NewCreateFilmsFromUrlsUseCase(loadFilmByUrlRepository dataprotocols.LoadFilmByUrlRepository) *CreateFilmsFromUrlsUseCase {
	return &CreateFilmsFromUrlsUseCase{
		loadFilmByUrlRepository: loadFilmByUrlRepository,
	}
}

func (usecase *CreateFilmsFromUrlsUseCase) Create(urls []string) (ids []string, err error) {
	for _, url := range urls {
		_, err := usecase.loadFilmByUrlRepository.Load(url)
		if err != nil {
			return nil, err
		}
	}
	return []string{}, nil
}
