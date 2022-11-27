package httpprotocolsstub

import "github.com/claudiomozer/starwarsapi/src/domain/usecases"

type LoadPlanetByIdRepositoryStub struct {
	Id          int
	TimesCalled int
}

func NewLoadPlanetByIdRepositoryStub() *LoadPlanetByIdRepositoryStub {
	return &LoadPlanetByIdRepositoryStub{}
}

func (stub *LoadPlanetByIdRepositoryStub) Load(id int) (*usecases.PlanetDTO, error) {
	stub.Id = id
	stub.TimesCalled++
	return nil, nil
}
