package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
)

type GetPlanetIdByUrlUseCaseSutParams struct {
	sut                        *data.GetPlanetIdByUrlUseCase
	getPlanetIdByUrlRepository *httpprotocolsstub.GetPlanetIdByUrlRepositoryStub
}

func MakeGetPlanetIdByUrlUseCaseSutParams() *GetPlanetIdByUrlUseCaseSutParams {
	getPlanetIdByUrlRepository := makeGetPlanetIdByUrlRepository()
	return &GetPlanetIdByUrlUseCaseSutParams{
		getPlanetIdByUrlRepository: getPlanetIdByUrlRepository,
		sut:                        data.NewGetPlanetIdByUrlUseCase(getPlanetIdByUrlRepository),
	}
}

func makeGetPlanetIdByUrlRepository() *httpprotocolsstub.GetPlanetIdByUrlRepositoryStub {
	return httpprotocolsstub.NewGetPlanetIdByUrlRepositoryStub()
}

func TestShouldCallGetPlanetIdByUrlRepositoryWithTheCorrectUrl(t *testing.T) {
	sutParams := MakeGetPlanetIdByUrlUseCaseSutParams()
	sut := sutParams.sut
	getPlanetIdByUrlUseCase := sutParams.getPlanetIdByUrlRepository
	urlTest := "url.test.com"
	sut.GetByUrl(urlTest)

	if getPlanetIdByUrlUseCase.TimesCalled == 0 {
		t.Error("Should call GetPlanetIdByUrlUseCaseRepository at least once")
	}

	if getPlanetIdByUrlUseCase.Url != urlTest {
		t.Errorf("Should call GetPlanetIdByUrlUseCaseRepository with: %s\n", urlTest)
	}
}
