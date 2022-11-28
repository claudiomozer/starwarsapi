package httprepo

import (
	"fmt"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
)

type PlanetRepository struct {
	getByIdUrl string
}

func NewPlanetRepository() *PlanetRepository {
	return &PlanetRepository{
		getByIdUrl: "https://swapi.dev/api/planets/",
	}
}

func (repo *PlanetRepository) Load(id int) (*domaindto.PlanetDTO, error) {
	status, _, err := Get(repo.getByIdUrl)

	if status == 404 {
		return nil, nil
	}

	return nil, err
}

func (repo *PlanetRepository) buildGetByIdURL(id int) string {
	return fmt.Sprintf("%s%d/", repo.getByIdUrl, id)
}
