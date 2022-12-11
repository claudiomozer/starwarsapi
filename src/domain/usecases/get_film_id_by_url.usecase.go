package usecases

type GetFilmIdBYUrlUseCase interface {
	GetByUrl(url string) (string, error)
}
