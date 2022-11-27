package httpprotocolsstub

import (
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/domain/mocks"
)

type LoadPlanetByIdRepositoryStub struct {
	Id          int
	TimesCalled int
	ReturnError bool
	ReturnsNil  bool
}

func NewLoadPlanetByIdRepositoryStub() *LoadPlanetByIdRepositoryStub {
	return &LoadPlanetByIdRepositoryStub{}
}

func (stub *LoadPlanetByIdRepositoryStub) Load(id int) (*domaindto.PlanetDTO, error) {
	stub.Id = id
	stub.TimesCalled++

	if stub.ReturnError {
		return nil, errors.New("error")
	}

	if stub.ReturnsNil {
		return nil, nil
	}

	return domainmocks.MockPlanetDTO(), nil
}
