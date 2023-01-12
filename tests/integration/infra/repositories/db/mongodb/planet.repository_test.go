package mongodbtest

import (
	"testing"

	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	"github.com/claudiomozer/starwarsapi/src/infra/repositories/db/mongodb"
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
}
