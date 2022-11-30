package httpprotocolsstub

import (
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type LoadFilmsByUrlRepositoryStub struct {
	Urls          []string
	TimesCalled   int
	ReturnErrorAt int
}

func NewLoadFilmsByUrlRepositoryStub() *LoadFilmsByUrlRepositoryStub {
	return &LoadFilmsByUrlRepositoryStub{}
}

func (stub *LoadFilmsByUrlRepositoryStub) Load(url string) ([]domaindto.FilmDTO, error) {
	stub.TimesCalled++

	if stub.ReturnErrorAt == stub.TimesCalled {
		return nil, errors.New("error")
	}

	stub.Urls = append(stub.Urls, url)
	return nil, nil
}
