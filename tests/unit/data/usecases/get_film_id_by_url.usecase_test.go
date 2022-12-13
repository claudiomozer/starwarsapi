package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
)

type GetFilmIdByUrlUseCaseSutParams struct {
	sut                      *data.GetFilmIdByUrlUseCase
	getFilmIdByUrlRepository *httpprotocolsstub.GetFilmIdByUrlRepositoryStub
}

func MakeGetFilmIdByUrlUseCaseSutParams() *GetFilmIdByUrlUseCaseSutParams {
	getFilmIdByUrlRepository := makeGetFilmIdByUrlRepository()
	return &GetFilmIdByUrlUseCaseSutParams{
		getFilmIdByUrlRepository: getFilmIdByUrlRepository,
		sut:                      data.NewGetFilmIdByUrlUseCase(getFilmIdByUrlRepository),
	}
}

func makeGetFilmIdByUrlRepository() *httpprotocolsstub.GetFilmIdByUrlRepositoryStub {
	return httpprotocolsstub.NewGetFilmIdByUrlRepositoryStub()
}

func TestShouldCallGetFilmIdByUrlUseCaseWithTheCorrectUrl(t *testing.T) {
	sutParams := MakeGetFilmIdByUrlUseCaseSutParams()
	sut := sutParams.sut
	getFilmIdByUrlUseCase := sutParams.getFilmIdByUrlRepository
	urlTest := "url.test.com"
	sut.GetByUrl(urlTest)

	if getFilmIdByUrlUseCase.TimesCalled == 0 {
		t.Error("Should call GetFilmIdByUrlUseCaseRepository at least once")
	}

	if getFilmIdByUrlUseCase.Url != urlTest {
		t.Errorf("Should call GetFilmIdByUrlUseCaseRepository with: %s\n", urlTest)
	}
}
