package mongodbtest

import (
	"testing"

	"github.com/benweissmann/memongo"
	domaindto "github.com/claudiomozer/starwarsapi/src/domain/dtos"
	"github.com/claudiomozer/starwarsapi/src/infra/repositories/db/mongodb"
)

var memServer *memongo.Server

func connectWithMongoServer() error {
	memServer, err := SetupMemoryServer()

	if err != nil {
		return err
	}

	sut := mongodb.NewMongoHelper("", "", "", "", "")
	sut.SetConnectionUri(memServer.URIWithRandomDB())
	err = sut.Connect()

	return err
}

func TestShouldReturnAnErrorIfThereIsNoConnection(t *testing.T) {
	sut := mongodb.NewFilmRepository()
	_, err := sut.Create(domaindto.NewFilmDTO("", "", "", ""))

	if err == nil {
		t.Error("Should return an error if there is no connections")
	}
}
