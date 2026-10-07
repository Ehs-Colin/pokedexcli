package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// GetPokemon
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// 1> Build the url to retreive
	url, err := url.JoinPath(baseURL, "pokemon", pokemonName)
	if err != nil {
		return Pokemon{}, err
	}
	// 2> Check CACHE first. Return if found in cache
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err = json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	// 3> Create a request for specified pokemonName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	// 4> Read the response from server int dat
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	// 5> Unmarshal response into go struct
	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	// 6> Add the url to cache
	c.cache.Add(url, dat)
	return pokemonResp, nil
}
