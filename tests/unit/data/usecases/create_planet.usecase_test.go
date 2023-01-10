package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

type createPlanetSutParam struct {
	sut                       *data.CreatePlanetUseCase
	createPlanetSutRepository *httpprotocolsstub.CreatePlanetRepositoryStub
}

func newCreatePlanetSutParams() *createPlanetSutParam {
	createPlanetSutRepository := makeCreatePlanetSutRepository()
	return &createPlanetSutParam{
		createPlanetSutRepository: createPlanetSutRepository,
		sut:                       data.NewCreatePlanetUseCase(createPlanetSutRepository),
	}
}

func makeCreatePlanetSutRepository() *httpprotocolsstub.CreatePlanetRepositoryStub {
	return httpprotocolsstub.NewCreatePlanetRepositoryStub()
}

func TestShouldCallCreatePlanetRepository(t *testing.T) {
	sutParams := newCreatePlanetSutParams()
	sut := sutParams.sut
	createPlanetSutRepository := sutParams.createPlanetSutRepository

	sut.Create(domainmocks.MockPlanetDbDTO())

	if createPlanetSutRepository.TimesCalled == 0 {
		t.Error("Should call CreatePlanetRepository at least once")
	}
}
