package httpdata

import (
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
	return usecase.repository.Load(id)
}
