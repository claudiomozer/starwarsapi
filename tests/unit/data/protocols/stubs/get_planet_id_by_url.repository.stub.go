package httpprotocolsstub

import (
	"errors"
)

type GetPlanetIdByUrlRepositoryStub struct {
	Url          string
	TimesCalled  int
	ReturnsError bool
	ReturnsEmpty bool
}

func NewGetPlanetIdByUrlRepositoryStub() *GetPlanetIdByUrlRepositoryStub {
	return &GetPlanetIdByUrlRepositoryStub{}
}

func (stub *GetPlanetIdByUrlRepositoryStub) GetByUrl(url string) (string, error) {
	stub.TimesCalled++
	stub.Url = url

	if stub.ReturnsError {
		return "", errors.New("error")
	}

	if stub.ReturnsEmpty {
		return "", nil
	}

	return "ObjectId(\"planet\")", nil
}
