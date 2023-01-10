package dataprotocols

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type CreatePlanetRepository interface {
	Create(planetDto *domaindto.PlanetDTO) (string, error)
}
