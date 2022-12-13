package httpprotocolsstub

import (
	"errors"
)

type GetFilmIdByUrlRepositoryStub struct {
	Url          string
	TimesCalled  int
	ReturnsError bool
	ReturnsEmpty bool
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

	if stub.ReturnsEmpty {
		return "", nil
	}

	return "ObjectId(\"teste\")", nil
}
