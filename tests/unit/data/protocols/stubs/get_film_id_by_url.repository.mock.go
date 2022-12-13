package httpprotocolsstub

import (
	"errors"
)

type GetFilmIdByUrlRepositoryStub struct {
	Url            string
	TimesCalled    int
	ReturnErrorAt  int
	ReturnsEmptyAt int
}

func NewGetFilmIdByUrlRepositoryStub() *GetFilmIdByUrlRepositoryStub {
	return &GetFilmIdByUrlRepositoryStub{}
}

func (stub *GetFilmIdByUrlRepositoryStub) GetByUrl(url string) (string, error) {
	stub.TimesCalled++
	stub.Url = url

	if stub.ReturnErrorAt == stub.TimesCalled {
		return "", errors.New("error")
	}

	if stub.ReturnsEmptyAt == stub.TimesCalled {
		return "", nil
	}

	return "ObjectId(\"teste\")", nil
}
