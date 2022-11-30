package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

type CreateFilmsFromUrlsSutParams struct {
	loadFilmsByUrlRepository *httpprotocolsstub.LoadFilmsByUrlRepositoryStub
	sut                      *data.CreateFilmsFromUrlsUseCase
}

func makeCreateFilmsFromUrlsSutParams() *CreateFilmsFromUrlsSutParams {
	loadFilmsByUrlRepository := makeLoadFilmsByUrlRepository()
	return &CreateFilmsFromUrlsSutParams{
		loadFilmsByUrlRepository: loadFilmsByUrlRepository,
		sut:                      data.NewCreateFilmsFromUrlsUseCase(loadFilmsByUrlRepository),
	}
}

func makeLoadFilmsByUrlRepository() *httpprotocolsstub.LoadFilmsByUrlRepositoryStub {
	return httpprotocolsstub.NewLoadFilmsByUrlRepositoryStub()
}

func TestShouldCallLoadFilmsByUrlRepositoryWithCorrectURLs(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	loadRepository := sutParams.loadFilmsByUrlRepository
	filmsUrls := domainmocks.MockPlanetDTO().Films

	sut.Create(filmsUrls)

	if loadRepository.TimesCalled != len(filmsUrls) {
		t.Errorf("Should call LoadFilmsByUrlRepository %d times\n", len(filmsUrls))
	}

	for key, url := range filmsUrls {
		if url != loadRepository.Urls[key] {
			t.Errorf("Should receive %s as URL\n", url)
		}
	}

}

func TestShouldReturnErrorIfLoadFilmsByURLRepositoryReturnsError(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	loadRepository := sutParams.loadFilmsByUrlRepository
	filmsUrls := domainmocks.MockPlanetDTO().Films
	callsUntilError := 2
	loadRepository.ReturnErrorAt = callsUntilError

	_, err := sut.Create(filmsUrls)

	if loadRepository.TimesCalled != callsUntilError {
		t.Errorf("Should call LoadFilmsByURLRepository %d times until fail\n", callsUntilError)
	}

	if err == nil {
		t.Error("Should return an error when LoadFilmsByURLRepository returns error")
	}
}
