package httpdatatest

import (
	"testing"

	httpdata "github.com/claudiomozer/starwarsapi/src/data/http/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/data/http/protocols/stubs"
)

type LoadPlanetByIdSutParams struct {
	loadPlanetByIdRepositoryStub *httpprotocolsstub.LoadPlanetByIdRepositoryStub
	sut                          *httpdata.LoadPlanetByIdUseCase
}

func makeLoadPlanetSutParams() *LoadPlanetByIdSutParams {
	loadPlanetByIdRepositoryStub := httpprotocolsstub.NewLoadPlanetByIdRepositoryStub()
	return &LoadPlanetByIdSutParams{
		loadPlanetByIdRepositoryStub: loadPlanetByIdRepositoryStub,
		sut:                          httpdata.NewLoadPlanetByIdUseCase(loadPlanetByIdRepositoryStub),
	}
}

func TestShouldCallLoadPlanetByIdRepositoryStubOnce(t *testing.T) {
	sutParams := makeLoadPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	sut := sutParams.loadPlanetByIdRepositoryStub
	sut.Load(1)

	if loadPlanetByIdRepositoryStub.TimesCalled == 0 {
		t.Error("Must call LoadPlanetByIdRepositoryStub at least once")
	}
}

func TestShouldReturnAnErrorIfLoadPlanetByIdRepositoryReturnsError(t *testing.T) {
	sutParams := makeLoadPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	loadPlanetByIdRepositoryStub.ReturnError = true
	sut := sutParams.sut
	planetDTO, err := sut.Load(1)

	if planetDTO != nil {
		t.Error("Should not return a planet if an error occours")
	}

	if err == nil {
		t.Error("Should return an error from repository ")
	}
}
