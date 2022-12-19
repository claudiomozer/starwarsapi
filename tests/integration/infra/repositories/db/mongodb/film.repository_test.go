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

	sut := mongodb.NewMongoHelper("dbtest", "", "", "", "")
	sut.SetConnectionUri(memServer.URI())
	err = sut.Connect()

	return err
}

func TestShouldReturnAnErrorIfThereIsNoConnection(t *testing.T) {
	sut := mongodb.NewFilmRepository()
	_, err := sut.Create(domaindto.NewFilmDTO("", "", "", ""))

	if err == nil {
		t.Error("Should return an error if there is no connections")
	}

	_, err = sut.GetByUrl("URL")

	if err == nil {
		t.Error("Should return an error if there is no connections")
	}
}

func TestShouldReturnsAnIdOnSuccess(t *testing.T) {

	if err := connectWithMongoServer(); err != nil {
		t.Errorf("Error while connecting to the database:%v\n", err)
		return
	}

	sut := mongodb.NewFilmRepository()

	id, err := sut.Create(domaindto.NewFilmDTO("Test Title", "Test Director", "1999-03-01", "https://testurl.com"))

	if err != nil {
		t.Error("Should not return an error on success")
	}

	if len(id) == 0 {
		return
	}
}
