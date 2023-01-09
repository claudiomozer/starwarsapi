package usecases

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type GetPlanetByApiIdUseCase interface {
	GetFromApi(id int) (*domaindto.PlanetDTO, error)
}
