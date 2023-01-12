package domaindto

type FilmDTO struct {
	Title       string `json:"title" bson:"title" `
	Director    string `json:"director" bson:"director"`
	ReleaseDate string `json:"release_date" bson:"release_date"`
	Url         string `json:"url" bson:"url"`
}

func NewFilmDTO(title, director, release_date, url string) *FilmDTO {
	return &FilmDTO{
		Title:       title,
		Director:    director,
		ReleaseDate: release_date,
		Url:         url,
	}
}
