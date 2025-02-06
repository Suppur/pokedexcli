package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchList(arg string) (Pokemons, error) {
	url := baseURL + "/pokemon/" + arg

	if val, ok := c.cache.Get(url); ok {
		pokeInfo := Pokemons{}
		err := json.Unmarshal(val, &pokeInfo)
		if err != nil {
			return Pokemons{}, err
		}
		return pokeInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemons{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemons{}, err
	}

	pokeL := Pokemons{}
	err = json.Unmarshal(data, &pokeL)
	if err != nil {
		return Pokemons{}, err
	}

	c.cache.Add(url, data)
	return pokeL, err
}
