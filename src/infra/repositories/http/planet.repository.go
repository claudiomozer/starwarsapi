package httprepo

import (
	"encoding/json"
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
	status, body, err := Get(repo.buildGetByIdURL(id))

	if status == 404 {
		return nil, nil
	}

	var planetDTO *domaindto.PlanetDTO = &domaindto.PlanetDTO{}
	err = json.Unmarshal(body, planetDTO)

	if err != nil {
		return nil, err
	}

	return planetDTO, err
}

func (repo *PlanetRepository) buildGetByIdURL(id int) string {
	return fmt.Sprintf("%s%d/", repo.getByIdUrl, id)
}
