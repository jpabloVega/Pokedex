package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationRes := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationRes)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationRes, nil
}
