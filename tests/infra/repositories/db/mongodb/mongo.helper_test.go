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
