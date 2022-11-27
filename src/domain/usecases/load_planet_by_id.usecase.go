package usecases

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type LoadPlanetByIdUseCase interface {
	Load(id int) (*domaindto.PlanetDTO, error)
}
