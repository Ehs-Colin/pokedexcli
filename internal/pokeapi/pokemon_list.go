package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const locationAreaPath = "location-area"

// ListLocations
func (c *Client) ListPokemon(locationName string) (LocationPokemon, error) {
	url, err := url.JoinPath(baseURL, locationAreaPath, locationName)
	if err != nil {
		return LocationPokemon{}, err
	}
	//Check CACHE first
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := LocationPokemon{}
		err = json.Unmarshal(val, pokemonResp)
		if err != nil {
			return LocationPokemon{}, err
		}
		return pokemonResp, nil
	}
	// Create a request for that locations pokemon
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemon{}, err
	}
	defer resp.Body.Close()
	// Read the response from server into dat
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationPokemon{}, err
	}
	// Unmarshal response into go struct
	locationsResp := LocationPokemon{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationPokemon{}, err
	}
	// Add the url to cache and return the results
	c.cache.Add(url, dat)
	return locationsResp, nil
}
