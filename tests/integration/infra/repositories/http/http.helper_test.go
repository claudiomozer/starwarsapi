package httptest

import (
	"testing"

	httpclient "github.com/claudiomozer/starwarsapi/src/infra/repositories/http"
)

func TestShouldReturnAResponseOnSuccess(t *testing.T) {
	status, body, err := httpclient.Get("https://swapi.dev/api/planets")

	if err != nil {
		t.Error("Should not return error on success")
	}

	if status != 200 {
		t.Error("Might return 200 for StatusCode")
	}

	if len(body) == 0 {
		t.Error("Should return a not empty body on success")
	}
}

func TestShouldReturnErrorIfInvalidUrlIsProvided(t *testing.T) {
	status, body, err := httpclient.Get("https://invalidurl.get/invalidresource")

	if err == nil {
		t.Error("Should return error when an invalid URL is given")
	}

	if status != 0 {
		t.Error("Should not return a valid status on error")
	}

	if len(body) != 0 {
		t.Error("Should return a empty body on error")
	}
}
