package httpprotocolsstub

import (
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

type GetPlanetByApiIdRepositoryStub struct {
	Id          int
	TimesCalled int
	ReturnError bool
	ReturnsNil  bool
}

func NewGetPlanetByApiIdRepositoryStub() *GetPlanetByApiIdRepositoryStub {
	return &GetPlanetByApiIdRepositoryStub{}
}

func (stub *GetPlanetByApiIdRepositoryStub) GetFromApi(id int) (*domaindto.PlanetDTO, error) {
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
