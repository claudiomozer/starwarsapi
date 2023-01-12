package mongodbtest

import (
	"testing"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	"github.com/claudiomozer/starwarsapi/src/infra/repositories/db/mongodb"
	domainmocks "github.com/claudiomozer/starwarsapi/tests/unit/domain/mocks"
)

func TestShouldReturnAnErrorIfThereIsNoConnetion(t *testing.T) {
	sut := mongodb.NewPlanetRepository()

	id, err := sut.Create(&domaindto.PlanetDTO{})

	if id != "" {
		t.Error("Should not return an id on error")
	}

	if err == nil {
		t.Error("Must return an error if there is no connection")
	}

	id, err = sut.GetByUrl("URL")

	if err == nil {
		t.Error("Should return an error if there is no connections")
	}

	if id != "" {
		t.Error("Should not return an id if there is no connection")
	}

}

func TestShouldCreateAPlanetOnSuccess(t *testing.T) {
	if err := ConnectWithMongoServer(); err != nil {
		t.Errorf("Error while connecting to the database:%v\n", err)
		return
	}

	filmRepo := mongodb.NewFilmRepository()
	filmMocked := domainmocks.MockANewHopeFilmDTO()

	id, err := filmRepo.Create(filmMocked)

	if err != nil || id == "" {
		t.Error("Could not create a planet without films")
		return
	}

	planetMocked := domainmocks.MockPlanetDbDTO()
	planetMocked.Films = []string{id}
	sut := mongodb.NewPlanetRepository()
	id, err = sut.Create(planetMocked)

	if id == "" {
		t.Error("Must return a valid id on success")
	}

	if err != nil {
		t.Error("May not return an error on success")
	}
}
