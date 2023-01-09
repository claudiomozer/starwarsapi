package data

import (
	"errors"
	"fmt"

	dataprotocols "github.com/claudiomozer/starwarsapi/src/data/protocols"
	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type GetPlanetByApiIdUseCase struct {
	repository dataprotocols.GetPlanetByApiIdRepository
}

func NewGetPlanetByApiIdUseCase(repository dataprotocols.GetPlanetByApiIdRepository) *GetPlanetByApiIdUseCase {
	return &GetPlanetByApiIdUseCase{
		repository: repository,
	}
}

func (usecase *GetPlanetByApiIdUseCase) GetFromApi(id int) (*domaindto.PlanetDTO, error) {
	planet, err := usecase.repository.GetFromApi(id)

	if planet == nil {
		return nil, errors.New(fmt.Sprintf("Nenhum planeta encontrado para o id: %d", id))
	}

	return planet, err
}
