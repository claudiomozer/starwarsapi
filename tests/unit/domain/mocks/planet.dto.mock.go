package domainmocks

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

func MockPlanetDTO() *domaindto.PlanetDTO {
	return domaindto.NewPlanetDTO(
		"Tatooine",
		"https://swapi.dev/api/planets/1/",
		"arid",
		"desert",
		[]string{
			"https://swapi.dev/api/films/1/",
			"https://swapi.dev/api/films/3/",
			"https://swapi.dev/api/films/4/",
			"https://swapi.dev/api/films/5/",
			"https://swapi.dev/api/films/6/",
		},
	)
}
