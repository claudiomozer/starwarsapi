package domaindto

type PlanetDTO struct {
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"films"`
}

func NewPlanetDTO(name, url, climate, terrain string, films []string) *PlanetDTO {
	return &PlanetDTO{
		Name:    name,
		Url:     url,
		Climate: climate,
		Terrain: terrain,
		Films:   films,
	}
}
