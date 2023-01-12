package mongodb

import (
	"context"
	"errors"
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	infrahelpers "github.com/claudiomozer/starwarsapi/src/infra/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlanetRepository struct {
	collection string
}

func NewPlanetRepository() *PlanetRepository {
	return &PlanetRepository{
		collection: "planets",
	}
}

type singleResultPlanetId struct {
	Id string `bson:"_id"`
}

func (repo *PlanetRepository) Create(planetDto *domaindto.PlanetDTO) (string, error) {

	if Helper == nil || (Helper != nil && Helper.IsConnectionInvalid()) {
		return "", errors.New("Erro ao criar planeta na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}

	collection := Helper.GetCollection(repo.collection)
	result, err := collection.InsertOne(context.TODO(), *planetDto)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Erro ao inserir Film na base de dados: %v\n", err))
	}

	objectId := result.InsertedID.(primitive.ObjectID)
	return objectId.String(), err
}

func (repo *PlanetRepository) GetByUrl(url string) (string, error) {
	if Helper == nil || (Helper != nil && Helper.IsConnectionInvalid()) {
		return "", errors.New("Erro ao buscar Film na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}

	if infrahelpers.IsUrlInvalid(url) {
		return "", errors.New("Impossível buscar filme na base de dados: URL inválida")
	}

	collection := Helper.GetCollection(repo.collection)
	singleResult := collection.FindOne(context.Background(), repo.buildGetByUrlFilter(url), repo.buildGetByUrl())
	singleResultId := &singleResultPlanetId{}

	if err := singleResult.Decode(singleResultId); err != nil {
		return "", err
	}

	return singleResultId.Id, nil
}

func (repo *PlanetRepository) buildGetByUrlFilter(url string) interface{} {
	return bson.D{primitive.E{Key: "url", Value: url}}
}

func (repo *PlanetRepository) buildGetByUrl() *options.FindOneOptions {
	return &options.FindOneOptions{
		Projection: bson.D{primitive.E{Key: "_id", Value: 1}},
	}
}
