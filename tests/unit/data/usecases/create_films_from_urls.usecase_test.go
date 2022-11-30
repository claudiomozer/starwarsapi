package httpdatatest

import (
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

type CreateFilmsFromUrlsSutParams struct {
	loadFilmByUrlRepository *httpprotocolsstub.LoadFilmByUrlRepositoryStub
	sut                     *data.CreateFilmsFromUrlsUseCase
}

func makeCreateFilmsFromUrlsSutParams() *CreateFilmsFromUrlsSutParams {
	loadFilmByUrlRepository := makeLoadFilmByUrlRepository()
	return &CreateFilmsFromUrlsSutParams{
		loadFilmByUrlRepository: loadFilmByUrlRepository,
		sut:                     data.NewCreateFilmsFromUrlsUseCase(loadFilmByUrlRepository),
	}
}

func makeLoadFilmByUrlRepository() *httpprotocolsstub.LoadFilmByUrlRepositoryStub {
	return httpprotocolsstub.NewLoadFilmByUrlRepositoryStub()
}

func TestShouldCallLoadFilmByUrlRepositoryWithCorrectURLs(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	loadRepository := sutParams.loadFilmByUrlRepository
	filmsUrls := domainmocks.MockPlanetDTO().Films

	sut.Create(filmsUrls)

	if loadRepository.TimesCalled != len(filmsUrls) {
		t.Errorf("Should call LoadFilmByUrlRepository %d times\n", len(filmsUrls))
	}

	for key, url := range filmsUrls {
		if url != loadRepository.Urls[key] {
			t.Errorf("Should receive %s as URL\n", url)
		}
	}

}

func TestShouldReturnErrorIfLoadFilmByURLRepositoryReturnsError(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	loadRepository := sutParams.loadFilmByUrlRepository
	filmsUrls := domainmocks.MockPlanetDTO().Films
	callsUntilError := 2
	loadRepository.ReturnErrorAt = callsUntilError

	_, err := sut.Create(filmsUrls)

	if loadRepository.TimesCalled != callsUntilError {
		t.Errorf("Should call LoadFilmByURLRepository %d times until fail\n", callsUntilError)
	}

	if err == nil {
		t.Error("Should return an error when LoadFilmByURLRepository returns error")
	}
}
