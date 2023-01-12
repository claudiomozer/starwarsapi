package mongodb

import (
	"errors"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type PlanetRepository struct {
	collection string
}

func NewPlanetRepository() *PlanetRepository {
	return &PlanetRepository{
		collection: "planets",
	}
}

func (repository *PlanetRepository) Create(planetDto *domaindto.PlanetDTO) (string, error) {

	if Helper == nil || (Helper != nil && Helper.IsConnectionInvalid()) {
		return "", errors.New("Erro ao criar planeta na base de dados. Nenhuma conexão com banco de dados estabelecida")
	}

	return "", nil
}
