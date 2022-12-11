package httptest

import (
	"testing"

	httprepo "github.com/claudiomozer/starwarsapi/src/infra/repositories/http"
)

func FuzzShouldReturnAnErrorIfAnInvalidURLIsGiven(f *testing.F) {
	f.Add("https://invalidurl.com.br/invalid/route")
	f.Add("https://swapi.dev/api/films/e")
	f.Add("https://swapi.dev/api/films/#")
	f.Add("http://swapi.dev/api/films/5")
	f.Add("https://swapi.dev/api/films/5e3e3")
	f.Add("https://swapi.dev/api/films/53^")
	f.Add("https://swapi.dev/api/films/a53")
	f.Add("http://swapi.dev/api/films/5/4")

	f.Fuzz(func(t *testing.T, invalidUrl string) {
		sut := httprepo.NewFilmRepository()
		_, err := sut.Load(invalidUrl)

		if err == nil {
			t.Error("Should return an error when an invalid URL is given")
		}
	})
}

func TestShouldReturnEmptyIfGivenUrlDoesNotExist(t *testing.T) {
	sut := httprepo.NewFilmRepository()
	film, err := sut.Load("https://swapi.dev/api/films/0")

	if err != nil {
		t.Error("Should not return an error if film does not exist")
	}

	if film != nil {
		t.Error("Film must be null")
	}
}

func TestShouldReturnAValidFilmDtoOnSuccess(t *testing.T) {
	sut := httprepo.NewFilmRepository()
	validUrl := "https://swapi.dev/api/films/1/"
	film, err := sut.Load(validUrl)

	if err != nil {
		t.Error("Should not return an error on success")
	}

	if film == nil {
		t.Error("Film must not be null on success")
		return
	}

	if film.Url != validUrl {
		t.Error("Must return a valid film on success")
	}

	if film.Title == "" {
		t.Error("Must return a film with a valid title")
	}

	if film.ReleaseDate == "" {
		t.Error("Must return a film with a valid release date")
	}

	if film.Director == "" {
		t.Error("Must return a film with a valid director")
	}
}
