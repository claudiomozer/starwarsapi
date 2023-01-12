package domaindto

type PlanetDTO struct {
	Name    string   `json:"name" bson:"name"`
	Url     string   `json:"url" bson:"url"`
	Climate string   `json:"climate" bson:"climate"`
	Terrain string   `json:"terrain" bson:"terrain"`
	Films   []string `json:"films" bson:"films"`
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
