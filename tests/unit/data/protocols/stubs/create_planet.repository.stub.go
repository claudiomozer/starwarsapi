package httpprotocolsstub

import (
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type CreatePlanetRepositoryStub struct {
	PlanetDTO   *domaindto.PlanetDTO
	TimesCalled int
}

func NewCreatePlanetRepositoryStub() *CreatePlanetRepositoryStub {
	return &CreatePlanetRepositoryStub{}
}

func (stub *CreatePlanetRepositoryStub) Create(planetDto *domaindto.PlanetDTO) (string, error) {
	stub.TimesCalled++
	return fmt.Sprintf("%d-id", stub.TimesCalled), nil
}
