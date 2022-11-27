package usecases

type PlanetDTO struct {
	name    string
	url     string
	climate string
	terrain string
	films   []string
}

type LoadPlanetByIdUseCase interface {
	Load(id int) (*PlanetDTO, error)
}
