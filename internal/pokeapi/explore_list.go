package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreList(arg string) (ExploreL, error) {
	url := baseURL + "/location-area/" + arg

	if val, ok := c.cache.Get(url); ok {
		locationResp := ExploreL{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return ExploreL{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreL{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreL{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreL{}, err
	}

	expl := ExploreL{}
	err = json.Unmarshal(data, &expl)
	if err != nil {
		return ExploreL{}, err
	}

	c.cache.Add(url, data)
	return expl, err
}
