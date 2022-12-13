package httpprotocolsstub

import (
	"errors"
)

type GetFilmIdByUrlRepositoryStub struct {
	Url          string
	TimesCalled  int
	ReturnsError bool
}

func NewGetFilmIdByUrlRepositoryStub() *GetFilmIdByUrlRepositoryStub {
	return &GetFilmIdByUrlRepositoryStub{}
}

func (stub *GetFilmIdByUrlRepositoryStub) GetByUrl(url string) (string, error) {
	stub.TimesCalled++
	stub.Url = url

	if stub.ReturnsError {
		return "", errors.New("error")
	}

	return "ObjectId(\"teste\")", nil
}
