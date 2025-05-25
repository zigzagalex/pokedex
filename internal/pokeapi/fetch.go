// internal/pokeapi/fetch.go
package pokeapi

import (
	"io"
	"net/http"
)

type LocationArea struct {
	Name string
	URL  string
}

type PokeAPIResult struct {
	Count    int
	Next     string
	Previous string
	Results  []LocationArea
}

func GetResult(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
