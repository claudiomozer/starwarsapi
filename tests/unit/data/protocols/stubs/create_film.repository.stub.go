package httpprotocolsstub

import (
	"errors"
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type CreateFilmRepositoryStub struct {
	FilmDTOs      []*domaindto.FilmDTO
	TimesCalled   int
	ReturnErrorAt int
}

func NewCreateFilmRepositoryStub() *CreateFilmRepositoryStub {
	return &CreateFilmRepositoryStub{}
}

func (stub *CreateFilmRepositoryStub) Create(filmDTO *domaindto.FilmDTO) (id string, err error) {
	stub.TimesCalled++

	if stub.TimesCalled == stub.ReturnErrorAt {
		id = ""
		err = errors.New("test")
		return
	}

	stub.FilmDTOs = append(stub.FilmDTOs, filmDTO)
	return fmt.Sprintf("%d-id", stub.TimesCalled), nil
}
