package models

import (
	"encoding/json"
	"spoty/internal/api"
	"strings"
)

type ResponseAlbum struct {
	Response ResultsAlbum `json:"results"`
}

type ResultsAlbum struct {
	Query struct {
		Text        string `json:"#text"`
		Role        string `json:"role"`
		SearchTerms string `json:"searchTerms"`
		StartPage   string `json:"startPage"`
	} `json:"opensearch:Query"`
	TotalRes     string       `json:"opensearch:totalResults"`
	StartIndex   string       `json:"opensearch:startIndex"`
	ItemsPerPage string       `json:"opensearch:itemsPerPage"`
	Matches      AlbumMatches `json:"albummatches"`
	Attributes   struct {
		For string `json:"for"`
	} `json:"@attr"`
}

type AlbumMatches struct {
	Albums []AlbumRes `json:"album"`
}

type AlbumRes struct {
	Name       string  `json:"name"`
	Artist     string  `json:"artist"`
	Url        string  `json:"url"`
	Images     []Image `json:"image"`
	Mbid       string  `json:"mbid"`
	Streamable string  `json:"streamable"`
}

func GetAlbumResults(query string) (ResultsAlbum, error) {
	var response ResponseAlbum

	request := strings.ReplaceAll(query, " ", "+")
	body, err := api.ParseJson("http://ws.audioscrobbler.com/2.0/?method=album.search&album=" + request + "&api_key=" + API_KEY + "&format=json")
	err = json.Unmarshal(body, &response)

	return response.Response, err
}
