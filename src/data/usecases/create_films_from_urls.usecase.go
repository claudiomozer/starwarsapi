package data

import dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"

type CreateFilmsFromUrlsUseCase struct {
	loadFilmsByUrlRepository dataprotocols.LoadFilmsByUrlRepository
}

func NewCreateFilmsFromUrlsUseCase(loadFilmsByUrlRepository dataprotocols.LoadFilmsByUrlRepository) *CreateFilmsFromUrlsUseCase {
	return &CreateFilmsFromUrlsUseCase{
		loadFilmsByUrlRepository: loadFilmsByUrlRepository,
	}
}

func (usecase *CreateFilmsFromUrlsUseCase) Create(urls []string) (ids []string, err error) {
	for _, url := range urls {
		_, err := usecase.loadFilmsByUrlRepository.Load(url)
		if err != nil {
			return nil, err
		}
	}
	return []string{}, nil
}
