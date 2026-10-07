package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// GetLocations
func (c *Client) ListPokemon(locationName string) (Location, error) {
	// 1> Build the correct url to retreive location area information
	url, err := url.JoinPath(baseURL, "location-area", locationName)
	if err != nil {
		return Location{}, err
	}
	// 2>Check CACHE first.  Return from cache if found
	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err = json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}
	// 3> Create a request for that locations pokemon
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()
	// 4> Read the response from server into dat
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	// 5> Unmarshal response into go struct
	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}
	// 6> Add the url to cache and return the results
	c.cache.Add(url, dat)
	return locationsResp, nil
}
