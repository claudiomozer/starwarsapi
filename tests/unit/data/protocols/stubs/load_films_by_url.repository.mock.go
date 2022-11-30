package httpprotocolsstub

import (
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type LoadFilmByUrlRepositoryStub struct {
	Urls          []string
	TimesCalled   int
	ReturnErrorAt int
}

func NewLoadFilmByUrlRepositoryStub() *LoadFilmByUrlRepositoryStub {
	return &LoadFilmByUrlRepositoryStub{}
}

func (stub *LoadFilmByUrlRepositoryStub) Load(url string) (*domaindto.FilmDTO, error) {
	stub.TimesCalled++

	if stub.ReturnErrorAt == stub.TimesCalled {
		return nil, errors.New("error")
	}

	stub.Urls = append(stub.Urls, url)
	return nil, nil
}
