// internal/pokeapi/fetch.go
package pokeapi

import (
	"encoding/json"
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

func GetLocationAreas(url string) (*PokeAPIResult, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data PokeAPIResult
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
