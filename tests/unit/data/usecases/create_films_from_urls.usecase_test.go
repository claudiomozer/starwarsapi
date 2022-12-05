package httpdatatest

import (
	"fmt"
	"testing"

	data "github.com/claudiomozer/starwarsapi/src/data/usecases"
	httpprotocolsstub "github.com/claudiomozer/starwarsapi/tests/unit/data/protocols/stubs"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

type CreateFilmsFromUrlsSutParams struct {
	loadFilmByUrlRepository *httpprotocolsstub.LoadFilmByUrlRepositoryStub
	createFilmRepository    *httpprotocolsstub.CreateFilmRepositoryStub
	sut                     *data.CreateFilmsFromUrlsUseCase
}

func makeCreateFilmsFromUrlsSutParams() *CreateFilmsFromUrlsSutParams {
	loadFilmByUrlRepository := makeLoadFilmByUrlRepository()
	createFilmRepository := makeCreateFilmRepository()
	return &CreateFilmsFromUrlsSutParams{
		loadFilmByUrlRepository: loadFilmByUrlRepository,
		createFilmRepository:    createFilmRepository,
		sut:                     data.NewCreateFilmsFromUrlsUseCase(loadFilmByUrlRepository, createFilmRepository),
	}
}

func makeLoadFilmByUrlRepository() *httpprotocolsstub.LoadFilmByUrlRepositoryStub {
	return httpprotocolsstub.NewLoadFilmByUrlRepositoryStub()
}

func makeCreateFilmRepository() *httpprotocolsstub.CreateFilmRepositoryStub {
	return httpprotocolsstub.NewCreateFilmRepositoryStub()
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

func TestShouldCallCreateFilmRepositoryWithCorrectValues(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	createRepository := sutParams.createFilmRepository
	filmsUrls := domainmocks.MockPlanetDTO().Films

	sut.Create(filmsUrls)

	if createRepository.TimesCalled == 0 {
		t.Error("Should call CreateFilmRepository at least once")
	}

	for k, film := range createRepository.FilmDTOs {
		if film.Url != filmsUrls[k] {
			t.Errorf("Should call CreateFilm repository with %s from URL %s", film.Url, filmsUrls[k])
		}
	}
}

func TestShouldReturnAnErrorIfCreateFilmRepositoryReturnsError(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	createRepository := sutParams.createFilmRepository
	callsUntilError := 2
	createRepository.ReturnErrorAt = callsUntilError

	_, err := sut.Create(domainmocks.MockPlanetDTO().Films)

	if createRepository.TimesCalled == 0 {
		t.Error("Should call CreateFilmRepository at least once")
	}

	if err == nil {
		t.Error("Should return an error if CreateFilmRepository returns error")
	}

	if createRepository.TimesCalled != callsUntilError {
		t.Errorf("Should call CreateFilmRepository %d times until fail\n", callsUntilError)
	}
}

func TestShouldReturnAnArrayOfIdsOnSuccess(t *testing.T) {
	sutParams := makeCreateFilmsFromUrlsSutParams()
	sut := sutParams.sut
	ids, err := sut.Create(domainmocks.MockPlanetDTO().Films)

	if err != nil {
		t.Error("Should not return an error on success")
	}

	for k, id := range ids {
		if id != fmt.Sprintf("%d-id", k+1) {
			t.Error("Should return the correct ids")
		}
	}
}
