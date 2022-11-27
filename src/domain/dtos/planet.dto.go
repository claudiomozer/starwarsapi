package domaindto

type PlanetDTO struct {
	name    string
	url     string
	climate string
	terrain string
	films   []string
}

func NewPlanetDTO(name, url, climate, terrain string, films []string) *PlanetDTO {
	return &PlanetDTO{
		name:    name,
		url:     url,
		climate: climate,
		terrain: terrain,
		films:   films,
	}
}

func (dto *PlanetDTO) GetURL() string {
	return dto.url
}
