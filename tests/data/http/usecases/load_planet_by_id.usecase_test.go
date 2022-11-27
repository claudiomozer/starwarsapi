package httpdatatest

import (
	"testing"

	httpdata "github.com/claudiomozer/starwarsapi/src/data/http/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/data/http/protocols/stubs"
)

func TestShouldCallLoadPlanetByIdRepositoryStubOnce(t *testing.T) {
	loadPlanetByIdRepositoryStub := httpprotocolsstub.NewLoadPlanetByIdRepositoryStub()
	sut := httpdata.NewLoadPlanetByIdUseCase(loadPlanetByIdRepositoryStub)
	sut.Load(1)

	if loadPlanetByIdRepositoryStub.TimesCalled == 0 {
		t.Error("Must call LoadPlanetByIdRepositoryStub at least once")
	}
}

func TestShouldReturnAnErrorIfLoadPlanetByIdRepositoryReturnsError(t *testing.T) {
	loadPlanetByIdRepositoryStub := httpprotocolsstub.NewLoadPlanetByIdRepositoryStub()
	loadPlanetByIdRepositoryStub.ReturnError = true
	sut := httpdata.NewLoadPlanetByIdUseCase(loadPlanetByIdRepositoryStub)
	planetDTO, err := sut.Load(1)

	if planetDTO != nil {
		t.Error("Should not return a planet if an error occours")
	}

	if err == nil {
		t.Error("Should return an error from repository ")
	}
}
