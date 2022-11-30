package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	protocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
)

type LoadPlanetByIdSutParams struct {
	loadPlanetByIdRepositoryStub *protocolsstub.LoadPlanetByIdRepositoryStub
	sut                          *data.LoadPlanetByIdUseCase
}

func makeLoadPlanetSutParams() *LoadPlanetByIdSutParams {
	loadPlanetByIdRepositoryStub := protocolsstub.NewLoadPlanetByIdRepositoryStub()
	return &LoadPlanetByIdSutParams{
		loadPlanetByIdRepositoryStub: loadPlanetByIdRepositoryStub,
		sut:                          data.NewLoadPlanetByIdUseCase(loadPlanetByIdRepositoryStub),
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

func TestShouldReturnAnErrorIfNilPlanetIsReturned(t *testing.T) {
	sutParams := makeLoadPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	loadPlanetByIdRepositoryStub.ReturnsNil = true
	sut := sutParams.sut
	planetDTO, err := sut.Load(1)

	if planetDTO != nil {
		t.Error("Should not return a planet if repository returns nils")
	}

	if err == nil {
		t.Error("Should return an error from repository")
	}
}

func TestShouldReturnAPlanetOnSuccess(t *testing.T) {
	sutParams := makeLoadPlanetSutParams()
	sut := sutParams.sut
	planetDTO, err := sut.Load(1)

	if err != nil {
		t.Error("Should not return an error on success")
	}

	if planetDTO == nil {
		t.Error("Should return a planet on success")
		return
	}

	if planetDTO.Url != "https://swapi.dev/api/planets/1/" {
		t.Error("Should return a valid planet")
	}
}