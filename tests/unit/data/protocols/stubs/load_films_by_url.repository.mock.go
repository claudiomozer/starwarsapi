package httpprotocolsstub

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type LoadFilmsByUrlRepositoryStub struct {
	Urls        []string
	TimesCalled int
}

func NewLoadFilmsByUrlRepositoryStub() *LoadFilmsByUrlRepositoryStub {
	return &LoadFilmsByUrlRepositoryStub{}
}

func (stub *LoadFilmsByUrlRepositoryStub) Load(url string) ([]domaindto.FilmDTO, error) {
	stub.TimesCalled++
	stub.Urls = append(stub.Urls, url)
	return nil, nil
}
