package mongodb

import (
	"context"
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type FilmRepository struct {
	collection string
}

func NewFilmRepository() *FilmRepository {
	return &FilmRepository{
		collection: "films",
	}
}

func (repo *FilmRepository) Create(filmDTO *domaindto.FilmDTO) (id string, err error) {
	if repo.isConnectionInvalid() {
		return "", errors.New("Erro ao criar Film na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}
	return "", nil
}

func (repo *FilmRepository) isConnectionInvalid() bool {
	if Helper == nil || (Helper != nil && Helper.GetCient() == nil) {
		return true
	}

	if err := Helper.GetCient().Ping(context.TODO(), nil); err != nil {
		return true
	}

	return false
}
