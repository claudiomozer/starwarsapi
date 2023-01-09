package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	protocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
)

type GetPlanetByApiIdSutParams struct {
	loadPlanetByIdRepositoryStub *protocolsstub.GetPlanetByApiIdRepositoryStub
	sut                          *data.GetPlanetByApiIdUseCase
}

func makeGetPlanetSutParams() *GetPlanetByApiIdSutParams {
	loadPlanetByIdRepositoryStub := protocolsstub.NewGetPlanetByApiIdRepositoryStub()
	return &GetPlanetByApiIdSutParams{
		loadPlanetByIdRepositoryStub: loadPlanetByIdRepositoryStub,
		sut:                          data.NewGetPlanetByApiIdUseCase(loadPlanetByIdRepositoryStub),
	}
}

func TestShouldCallGetPlanetByApiIdRepositoryStubOnce(t *testing.T) {
	sutParams := makeGetPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	sut := sutParams.loadPlanetByIdRepositoryStub
	sut.GetFromApi(1)

	if loadPlanetByIdRepositoryStub.TimesCalled == 0 {
		t.Error("Must call GetPlanetByApiIdRepositoryStub at least once")
	}
}

func TestShouldReturnAnErrorIfGetPlanetByApiIdRepositoryReturnsError(t *testing.T) {
	sutParams := makeGetPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	loadPlanetByIdRepositoryStub.ReturnError = true
	sut := sutParams.sut
	planetDTO, err := sut.GetFromApi(1)

	if planetDTO != nil {
		t.Error("Should not return a planet if an error occours")
	}

	if err == nil {
		t.Error("Should return an error from repository ")
	}
}

func TestShouldReturnAnErrorIfNilPlanetIsReturned(t *testing.T) {
	sutParams := makeGetPlanetSutParams()
	loadPlanetByIdRepositoryStub := sutParams.loadPlanetByIdRepositoryStub
	loadPlanetByIdRepositoryStub.ReturnsNil = true
	sut := sutParams.sut
	planetDTO, err := sut.GetFromApi(1)

	if planetDTO != nil {
		t.Error("Should not return a planet if repository returns nils")
	}

	if err == nil {
		t.Error("Should return an error from repository")
	}
}

func TestShouldReturnAPlanetOnSuccess(t *testing.T) {
	sutParams := makeGetPlanetSutParams()
	sut := sutParams.sut
	planetDTO, err := sut.GetFromApi(1)

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
