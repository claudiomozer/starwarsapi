package dataprotocols

type GetFilmIdByUrlRepository interface {
	GetByUrl(url string) (string, error)
}
