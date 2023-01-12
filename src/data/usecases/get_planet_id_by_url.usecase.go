package data

import dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"

type GetPlanetIdByUrlUseCase struct {
	getPlanetIdByUrlRepository dataprotocols.GetPlanetIdByUrlRepository
}

func NewGetPlanetIdByUrlUseCase(repo dataprotocols.GetPlanetIdByUrlRepository) *GetPlanetIdByUrlUseCase {
	return &GetPlanetIdByUrlUseCase{
		getPlanetIdByUrlRepository: repo,
	}
}

func (usecase *GetPlanetIdByUrlUseCase) GetByUrl(url string) (string, error) {
	return usecase.getPlanetIdByUrlRepository.GetByUrl(url)
}
