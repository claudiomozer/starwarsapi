package domaindto

type FilmDTO struct {
	Name        string
	Director    string
	ReleaseDate string
	Url         string
}

func NewFilmDTO(name, director, release_date, url string) *FilmDTO {
	return &FilmDTO{
		Name:        name,
		Director:    director,
		ReleaseDate: release_date,
		Url:         url,
	}
}
