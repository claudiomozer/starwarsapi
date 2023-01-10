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

func TestShouldReturnAnErrorIfCreatePlanetRepositoryAlsoReturns(t *testing.T) {
	sutParams := newCreatePlanetSutParams()
	sut := sutParams.sut
	sutParams.createPlanetSutRepository.ReturnsError = true

	id, err := sut.Create(domainmocks.MockPlanetDbDTO())

	if id != "" {
		t.Error("Should not return a valid id if error is returned")
	}

	if err == nil {
		t.Error("Must return an error if CreatePlanetRepository returns error")
	}
}

func TestShouldReturnAnIdOnSuccess(t *testing.T) {
	sutParams := newCreatePlanetSutParams()
	sut := sutParams.sut

	id, err := sut.Create(domainmocks.MockPlanetDbDTO())

	if id == "" {
		t.Error("Should return a valid id on success")
	}

	if err != nil {
		t.Error("Should not return an error on success")
	}
}
