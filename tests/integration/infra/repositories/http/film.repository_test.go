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

	f.Fuzz(func(t *testing.T, invalidUrl string) {
		sut := httprepo.NewFilmRepository()
		_, err := sut.Load(invalidUrl)

		if err == nil {
			t.Error("Should return an error when an invalid URL is given")
		}
	})
}
