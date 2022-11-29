package usecases

type CreateFilmsFromUrlsUseCase interface {
	Create(urls []string) (ids []string, err error)
}
