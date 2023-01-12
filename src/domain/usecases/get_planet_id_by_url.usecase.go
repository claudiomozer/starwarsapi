package usecases

type GetPlanetIdByUrlUseCase interface {
	GetByUrl(url string) (string, error)
}
