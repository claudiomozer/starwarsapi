package domainmocks

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

func MockANewHopeFilmDTO() *domaindto.FilmDTO {
	return domaindto.NewFilmDTO(
		"A New Hope",
		"George Lucas",
		"1977-05-25",
		"https://swapi.dev/api/films/1/",
	)
}

func MockReturnOfTheJediFilmDTO() *domaindto.FilmDTO {
	return domaindto.NewFilmDTO(
		"Return of the Jedi",
		"Richard Marquand",
		"1983-05-25",
		"https://swapi.dev/api/films/3/",
	)
}

func MockThePhantomMenaceFilmDTO() *domaindto.FilmDTO {
	return domaindto.NewFilmDTO(
		"The Phantom Menace",
		"George Lucas",
		"1999-05-19",
		"https://swapi.dev/api/films/4/",
	)
}

func MockAttackOfTheClonesFilmDTO() *domaindto.FilmDTO {
	return domaindto.NewFilmDTO(
		"Attack of the Clones",
		"George Lucas",
		"2002-05-16",
		"https://swapi.dev/api/films/5/",
	)
}

func MockRevengeOfTheSithFilmDTO() *domaindto.FilmDTO {
	return domaindto.NewFilmDTO(
		"Revenge of the Sith",
		"George Lucas",
		"2005-05-19",
		"https://swapi.dev/api/films/6/",
	)
}

func MockFilmsInTatooine() []*domaindto.FilmDTO {
	return []*domaindto.FilmDTO{
		MockANewHopeFilmDTO(),
		MockReturnOfTheJediFilmDTO(),
		MockThePhantomMenaceFilmDTO(),
		MockAttackOfTheClonesFilmDTO(),
		MockRevengeOfTheSithFilmDTO(),
	}
}
