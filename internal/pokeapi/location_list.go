package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/billLee3/pokedex/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	pokecache := pokecache.NewCache(30 * time.Second)
	_, ok := pokecache.CacheEntries[url]
	if ok {
		responseCache, _ := pokecache.Get(url)
		var response RespShallowLocations
		err := json.Unmarshal(responseCache, &response)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return response, nil
	}
	fmt.Printf("%v", pokecache)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
