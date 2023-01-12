package infrahelpers

import "regexp"

func IsUrlInvalid(url string) bool {
	regex := regexp.MustCompile(`https://swapi.dev/api/(films|planets)/\d+/{0,1}$`)
	return !regex.MatchString(url)
}
