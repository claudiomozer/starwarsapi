package dataprotocols

type GetPlanetIdByUrlRepository interface {
	GetByUrl(url string) (string, error)
}
