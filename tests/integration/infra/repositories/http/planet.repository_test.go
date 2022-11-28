package httptest

import (
	"testing"

	httprepo "github.com/claudiomozer/starwarsapi/src/infra/repositories/http"
)

func TestShouldReturnAnEmptyDTOIfAnInvalidIdIsGiven(t *testing.T) {
	sut := httprepo.NewPlanetRepository()
	planetDTO, err := sut.Load(0)

	if planetDTO != nil {
		t.Error("Should not return a planetDto if invalid Id is given")
	}

	if err != nil {
		t.Error("Should not return an error if an invalid id is given")
	}
}

func TestShouldReturnAnValidDTOonSuccess(t *testing.T) {
	sut := httprepo.NewPlanetRepository()
	planetDTO, err := sut.Load(1)

	if planetDTO == nil {
		t.Errorf("Should return a planet DTO on success")
		return
	}

	if len(planetDTO.Url) == 0 {
		t.Errorf("Should return a planet DTO with a valid URL ion success")
	}

	if err != nil {
		t.Error("Should not return an error on success")
	}
}
