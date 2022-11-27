package httpdata

import (
	"errors"
	"fmt"

	httpdataprotocols "github.com/claudiomozer/starwarsapi/src/data/http/protocols"
	"github.com/claudiomozer/starwarsapi/src/domain/usecases"
)

type LoadPlanetByIdUseCase struct {
	repository httpdataprotocols.LoadPlanetByIdRepository
}

func NewLoadPlanetByIdUseCase(repository httpdataprotocols.LoadPlanetByIdRepository) *LoadPlanetByIdUseCase {
	return &LoadPlanetByIdUseCase{
		repository: repository,
	}
}

func (usecase *LoadPlanetByIdUseCase) Load(id int) (*usecases.PlanetDTO, error) {
	planet, err := usecase.repository.Load(id)

	if planet == nil {
		return nil, errors.New(fmt.Sprintf("Nenhum planeta encontrado para o id: %d", id))
	}

	return planet, err
}
