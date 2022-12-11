package mongodbtest

import (
	"testing"

	"github.com/claudiomozer/starwarsapi/src/infra/repositories/db/mongodb"
)

func TestShouldReturnAnErrorIfEmptyVarsAreGiven(t *testing.T) {
	sut := mongodb.NewMongoHelper("", "", "", "", "")

	if err := sut.Connect(); err == nil {
		t.Error("Should return an error if empty vars are given")
	}
}

func TestShouldReturnErrorIfConnectionFails(t *testing.T) {
	sut := mongodb.NewMongoHelper("INVALID_DATABASE", "INVALID_HOST", "INVALID_PORT", "INVALID_USER", "INVALID_PASS")

	if err := sut.Connect(); err == nil {
		t.Error("Should return an error if connection fails")
	}
}

func TestShouldConnectOnSuccess(t *testing.T) {
	memServer, err := SetupMemoryServer()

	if err != nil {
		t.Errorf("Error on setup server in memory:%v\n", err)
		return
	}
	defer memServer.Stop()

	sut := mongodb.NewMongoHelper("", "", "", "", "")
	sut.SetConnectionUri(memServer.URIWithRandomDB())

	if err := sut.Connect(); err != nil {
		t.Error("Should return an error if connection fails")
	}

	if sut.GetCient() == nil {
		t.Error("Should return a client on success")
	}

	sut.Disconnect()
}

func TestShouldReturnACollectionOnSuccess(t *testing.T) {
	memServer, err := SetupMemoryServer()

	if err != nil {
		t.Errorf("Error on setup server in memory:%v\n", err)
		return
	}
	defer memServer.Stop()

	sut := mongodb.NewMongoHelper("", "", "", "", "")
	sut.SetConnectionUri(memServer.URIWithRandomDB())
	err = sut.Connect()

	if err != nil {
		t.Errorf("Error while connecting with memory server: %v\n", err)
	}

	coll := sut.GetCollection("teste")

	if coll == nil {
		t.Error("Should return a valid collection on success")
	}
}
