package mongodb

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	collection := Helper.GetCollection(repo.collection)
	result, err := collection.InsertOne(context.TODO(), repo.getBson(filmDTO))

	if err != nil {
		return "", errors.New(fmt.Sprintf("Erro ao inserir Film na base de dados: %v\n", err))
	}

	objectId := result.InsertedID.(primitive.ObjectID)
	return objectId.String(), err
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

func (repo *FilmRepository) getBson(filmDTO *domaindto.FilmDTO) interface{} {
	return bson.D{
		primitive.E{Key: "title", Value: filmDTO.Title},
		primitive.E{Key: "director", Value: filmDTO.Director},
		primitive.E{Key: "release_date", Value: filmDTO.ReleaseDate},
		primitive.E{Key: "url", Value: filmDTO.Url},
	}
}

func (repo *FilmRepository) GetByUrl(url string) (string, error) {
	if repo.isConnectionInvalid() {
		return "", errors.New("Erro ao criar Film na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}

	if repo.isUrlInvalid(url) {
		return "", errors.New("Impossível buscar filme na base de dados: URL inválida")
	}

	return "", nil
}

func (repo *FilmRepository) isUrlInvalid(url string) bool {
	regex := regexp.MustCompile(`https://swapi.dev/api/films/\d+/{0,1}$`)
	return !regex.MatchString(url)
}
