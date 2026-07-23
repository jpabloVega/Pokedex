package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListStats(pokeName string) (PokeStats, error) {
	url := baseURL + "/pokemon/" + pokeName

	if val, ok := c.cache.Get(url); ok {
		statsRes := PokeStats{}
		err := json.Unmarshal(val, &statsRes)
		if err != nil {
			return PokeStats{}, err
		}
		return statsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeStats{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeStats{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		errorSC := errors.New(res.Status)
		return PokeStats{}, errorSC
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeStats{}, err
	}

	resStats := PokeStats{}
	err = json.Unmarshal(data, &resStats)
	if err != nil {
		return PokeStats{}, err
	}

	c.cache.Add(url, data)
	return resStats, nil
}
