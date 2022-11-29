package data

import (
	"errors"
	"fmt"

	dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"
	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type LoadPlanetByIdUseCase struct {
	repository dataprotocols.LoadPlanetByIdRepository
}

func NewLoadPlanetByIdUseCase(repository dataprotocols.LoadPlanetByIdRepository) *LoadPlanetByIdUseCase {
	return &LoadPlanetByIdUseCase{
		repository: repository,
	}
}

func (usecase *LoadPlanetByIdUseCase) Load(id int) (*domaindto.PlanetDTO, error) {
	planet, err := usecase.repository.Load(id)

	if planet == nil {
		return nil, errors.New(fmt.Sprintf("Nenhum planeta encontrado para o id: %d", id))
	}

	return planet, err
}
