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

func TestShouldCallGetFilmIdByUrlRepositoryWithTheCorrectUrl(t *testing.T) {
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

func TestShouldReturnAnErrorIfGetFilmIdByUrlRepositoryReturnsError(t *testing.T) {
	sutParams := MakeGetFilmIdByUrlUseCaseSutParams()
	sut := sutParams.sut
	sutParams.getFilmIdByUrlRepository.ReturnsError = true
	id, err := sut.GetByUrl("url.test.com")

	if id != "" {
		t.Error("Should not return an valid id if error is returned")
	}

	if err == nil {
		t.Error("Should return an error if GetFilmIdByUrlRepository returns error")
	}
}

func TestShouldReturnAnEmptyIfGetFilmIdByUrlRepositoryReturnsEmpty(t *testing.T) {
	sutParams := MakeGetFilmIdByUrlUseCaseSutParams()
	sut := sutParams.sut
	sutParams.getFilmIdByUrlRepository.ReturnsEmpty = true
	id, err := sut.GetByUrl("url.test.com")

	if id != "" {
		t.Error("Should not return an valid id if empty is returned")
	}

	if err != nil {
		t.Error("Should not return an error empty is returned")
	}
}

func TestShouldReturnAValidIdOnSuccess(t *testing.T) {
	sutParams := MakeGetFilmIdByUrlUseCaseSutParams()
	sut := sutParams.sut

	id, err := sut.GetByUrl("url.test.com")
	if id == "" {
		t.Error("Should return a valid id on success")
	}

	if err != nil {
		t.Error("Should not return an error on success")
	}
}
