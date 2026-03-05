package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

const (
	defaultURL string = "https://pokeapi.co/api/v2/"
)

// Client handles requests to the PokeAPI
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient creates a new PokeAPI client with a specific timeout
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}

func (c *Client) ListLocations(pageUrl *string) (*APIPokeReponseLocation, error) {
	baseURL := defaultURL + "location-area"
	if pageUrl != nil {
		baseURL = *pageUrl
	}

	if cacheResponse, ok := c.cache.Get(baseURL); ok {
		var response APIPokeReponseLocation
		if err := json.Unmarshal(cacheResponse, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	// Get the location list from the given URL
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get Data stream
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a slice of strings
	var response APIPokeReponseLocation
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Cache the response for future requests
	c.cache.Add(baseURL, data)

	return &response, nil
}

func (c *Client) ExploreArea(area string) (*APIPokeResponseArea, error) {
	baseURL := defaultURL + "location-area/" + area

	if cacheResponse, ok := c.cache.Get(baseURL); ok {
		var response APIPokeResponseArea
		if err := json.Unmarshal(cacheResponse, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	// Get the location list from the given URL
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get Data stream
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a slice of strings
	var response APIPokeResponseArea
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Cache the response for future requests
	c.cache.Add(baseURL, data)

	return &response, nil
}

func (c *Client) GetPokemon(pokemon string) (*Pokemon, error) {
	baseURL := defaultURL + "pokemon/" + pokemon

	if cacheResponse, ok := c.cache.Get(baseURL); ok {
		var response Pokemon
		if err := json.Unmarshal(cacheResponse, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	// Get the location list from the given URL
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get Data stream
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a slice of strings
	var response Pokemon
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Cache the response for future requests
	c.cache.Add(baseURL, data)

	return &response, nil
}
