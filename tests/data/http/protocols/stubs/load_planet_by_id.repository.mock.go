package httpprotocolsstub

import (
	"errors"

	"github.com/claudiomozer/starwarsapi/src/domain/usecases"
)

type LoadPlanetByIdRepositoryStub struct {
	Id          int
	TimesCalled int
	ReturnError bool
}

func NewLoadPlanetByIdRepositoryStub() *LoadPlanetByIdRepositoryStub {
	return &LoadPlanetByIdRepositoryStub{}
}

func (stub *LoadPlanetByIdRepositoryStub) Load(id int) (*usecases.PlanetDTO, error) {
	stub.Id = id
	stub.TimesCalled++

	if stub.ReturnError {
		return nil, errors.New("error")
	}

	return nil, nil
}
