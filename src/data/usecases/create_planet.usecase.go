package data

import (
	dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"
	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type CreatePlanetUseCase struct {
	createPlanetRepository dataprotocols.CreatePlanetRepository
}

func NewCreatePlanetUseCase(createPlanetRepository dataprotocols.CreatePlanetRepository) *CreatePlanetUseCase {
	return &CreatePlanetUseCase{
		createPlanetRepository: createPlanetRepository,
	}
}

func (usecase *CreatePlanetUseCase) Create(planetDto *domaindto.PlanetDTO) (string, error) {
	return usecase.createPlanetRepository.Create(planetDto)
}
