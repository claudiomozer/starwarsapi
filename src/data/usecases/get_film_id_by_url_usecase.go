package data

import dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"

type GetFilmIdByUrlUseCase struct {
	getFilmIdByUrlRepository dataprotocols.GetFilmIdByUrlRepository
}

func NewGetFilmIdByUrlUseCase(getFilmIdByUrlRepository dataprotocols.GetFilmIdByUrlRepository) *GetFilmIdByUrlUseCase {
	return &GetFilmIdByUrlUseCase{
		getFilmIdByUrlRepository: getFilmIdByUrlRepository,
	}
}

func (usecase *GetFilmIdByUrlUseCase) GetByUrl(url string) (string, error) {
	_, err := usecase.getFilmIdByUrlRepository.GetByUrl(url)
	return "", err
}
