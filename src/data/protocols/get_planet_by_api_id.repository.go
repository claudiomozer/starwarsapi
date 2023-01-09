package dataprotocols

import domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"

type GetPlanetByApiIdRepository interface {
	GetFromApi(id int) (*domaindto.PlanetDTO, error)
}
