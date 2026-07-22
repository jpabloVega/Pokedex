package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) (PokeData, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		encounterRes := PokeData{}
		err := json.Unmarshal(val, &encounterRes)
		if err != nil {
			return PokeData{}, err
		}
		return encounterRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeData{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		errorSC := errors.New(res.Status)
		return PokeData{}, errorSC
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeData{}, err
	}

	pokemonRes := PokeData{}
	err = json.Unmarshal(dat, &pokemonRes)
	if err != nil {
		return PokeData{}, err
	}

	c.cache.Add(url, dat)
	return pokemonRes, nil
}
