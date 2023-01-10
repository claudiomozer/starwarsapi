package httpprotocolsstub

import (
	"errors"
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type CreatePlanetRepositoryStub struct {
	PlanetDTO    *domaindto.PlanetDTO
	TimesCalled  int
	ReturnsError bool
}

func NewCreatePlanetRepositoryStub() *CreatePlanetRepositoryStub {
	return &CreatePlanetRepositoryStub{}
}

func (stub *CreatePlanetRepositoryStub) Create(planetDto *domaindto.PlanetDTO) (string, error) {
	stub.TimesCalled++

	if stub.ReturnsError {
		return "", errors.New("error")
	}

	return fmt.Sprintf("%d-id", stub.TimesCalled), nil
}
