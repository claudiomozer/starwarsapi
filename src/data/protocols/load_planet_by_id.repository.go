package dataprotocols

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type LoadPlanetByIdRepository interface {
	Load(id int) (*domaindto.PlanetDTO, error)
}
