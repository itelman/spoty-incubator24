package models

import (
	"encoding/json"
	"spoty/internal/api"
	"strings"
)

type TopAlbums struct {
	TopAlbums TopAlbum `json:"topalbums"`
}

type TopAlbum struct {
	Albums     []Album `json:"album"`
	Attributes struct {
		Artist     string `json:"artist"`
		Page       string `json:"page"`
		PerPage    string `json:"perPage"`
		TotalPages string `json:"totalPages"`
		Total      string `json:"total"`
	} `json:"@attr"`
}

type Album struct {
	Name      string          `json:"name"`
	PlayCount int             `json:"playcount"`
	Url       string          `json:"url"`
	Artist    ArtistShortInfo `json:"artist"`
	Images    []Image         `json:"image"`
}

type ArtistShortInfo struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	Url  string `json:"url"`
}

func GetArtistsTopAlbums(artistName string) (TopAlbum, error) {
	var topAlbums TopAlbums

	request := strings.ReplaceAll(artistName, " ", "+")
	body, err := api.ParseJson("http://ws.audioscrobbler.com/2.0/?method=artist.gettopalbums&artist=" + request + "&api_key=" + API_KEY + "&format=json")
	err = json.Unmarshal(body, &topAlbums)

	return topAlbums.TopAlbums, err
}
