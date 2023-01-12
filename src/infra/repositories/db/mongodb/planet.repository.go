package mongodb

import (
	"context"
	"errors"
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetRepository struct {
	collection string
}

func NewPlanetRepository() *PlanetRepository {
	return &PlanetRepository{
		collection: "planets",
	}
}

func (repo *PlanetRepository) Create(planetDto *domaindto.PlanetDTO) (string, error) {

	if Helper == nil || (Helper != nil && Helper.IsConnectionInvalid()) {
		return "", errors.New("Erro ao criar planeta na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}

	collection := Helper.GetCollection(repo.collection)
	result, err := collection.InsertOne(context.TODO(), repo.getBson(planetDto))

	if err != nil {
		return "", errors.New(fmt.Sprintf("Erro ao inserir Film na base de dados: %v\n", err))
	}

	objectId := result.InsertedID.(primitive.ObjectID)
	return objectId.String(), err
}

func (repo *PlanetRepository) getBson(planetDTO *domaindto.PlanetDTO) interface{} {
	return bson.D{
		primitive.E{Key: "name", Value: planetDTO.Name},
		primitive.E{Key: "climate", Value: planetDTO.Climate},
		primitive.E{Key: "terrain", Value: planetDTO.Terrain},
		primitive.E{Key: "films", Value: planetDTO.Films},
		primitive.E{Key: "url", Value: planetDTO.Url},
	}
}
