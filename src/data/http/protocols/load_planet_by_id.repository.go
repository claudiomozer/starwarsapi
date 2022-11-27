package httpdataprotocols

import "github.com/claudiomozer/starwarsapi/src/domain/usecases"

type LoadPlanetByIdRepository interface {
	Load(id int) (*usecases.PlanetDTO, error)
}
