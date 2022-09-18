package anilist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/metafates/mangal/log"
	"net/http"
	"strconv"
)

type anilistResponse struct {
	Data struct {
		Page struct {
			Media []*Manga `json:"media"`
		} `json:"page"`
	} `json:"data"`
}

var searchCache = make(map[string][]*Manga)

func Search(name string) ([]*Manga, error) {
	if mangas, ok := searchCache[name]; ok {
		return mangas, nil
	}

	// prepare body
	log.Info("Searching anilist for manga: " + name)
	body := map[string]interface{}{
		"query": searchQuery,
		"variables": map[string]interface{}{
			"query": name,
		},
	}

	// parse body to json
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// send request
	log.Info("Sending request to Anilist")
	resp, err := http.Post(
		"https://graphql.anilist.co",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("Anilist returned status code " + strconv.Itoa(resp.StatusCode))
		return nil, fmt.Errorf("invalid response code %d", resp.StatusCode)
	}

	// decode response
	var response anilistResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Error(err)
		return nil, err
	}

	mangas := response.Data.Page.Media
	log.Info("Got response from Anilist, found " + strconv.Itoa(len(mangas)) + " results")
	searchCache[name] = mangas
	return mangas, nil
}
