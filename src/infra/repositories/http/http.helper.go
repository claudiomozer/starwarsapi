package httprepo

import (
	"io"
	"net/http"
)

func Get(url string) (status int, body []byte, err error) {

	response, err := http.Get(url)

	if err == nil {
		defer response.Body.Close()
		status = response.StatusCode
		body, err = io.ReadAll(response.Body)
	}

	return
}
