package infrahelperstest

import (
	"testing"

	infrahelpers "github.com/claudiomozer/starwarsapi/src/infra/helpers"
)

func FuzzShoulReturnTrueIfAnInvalidURLIsGiven(f *testing.F) {
	f.Add("https://invalidurl.com.br/invalid/route")
	f.Add("https://swapi.dev/api/films/e")
	f.Add("https://swapi.dev/api/films/#")
	f.Add("http://swapi.dev/api/films/5")
	f.Add("https://swapi.dev/api/films/5e3e3")
	f.Add("https://swapi.dev/api/films/53^")
	f.Add("https://swapi.dev/api/films/a53")
	f.Add("http://swapi.dev/api/films/5/4")
	f.Add("https://swapi.dev/api/planets/e")
	f.Add("https://swapi.dev/api/planets/#")
	f.Add("http://swapi.dev/api/planets/1")
	f.Add("https://swapi.dev/api/planets/24e3e3")
	f.Add("https://swapi.dev/api/planets/12^")
	f.Add("https://swapi.dev/api/planets/a34")
	f.Add("http://swapi.dev/api/planets/2/1")

	f.Fuzz(func(t *testing.T, invalidUrl string) {
		isInvalid := infrahelpers.IsUrlInvalid(invalidUrl)

		if !isInvalid {
			t.Error("Should return true when an invalid URL is given")
		}
	})
}

func FuzzShouldReturnFalseIfUrlIsValid(f *testing.F) {
	f.Add("https://swapi.dev/api/films/5")
	f.Add("https://swapi.dev/api/films/53")
	f.Add("https://swapi.dev/api/films/3/")
	f.Add("https://swapi.dev/api/planets/1")
	f.Add("https://swapi.dev/api/planets/22")
	f.Add("https://swapi.dev/api/planets/4/")

	f.Fuzz(func(t *testing.T, invalidUrl string) {
		isInvalid := infrahelpers.IsUrlInvalid(invalidUrl)

		if isInvalid {
			t.Error("Should return true when an valid URL is given")
		}
	})
}
