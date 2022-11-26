package usecases

type PlanetDTO struct {
	name    string
	climate string
	terrain string
	films   []string
}

type LoadPlanetByIdUseCase interface {
	load(id int) *PlanetDTO
}
