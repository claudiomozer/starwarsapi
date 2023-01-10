package usecases

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type CreatePlanetUseCase interface {
	Create(planetDto *domaindto.PlanetDTO) (string, error)
}
