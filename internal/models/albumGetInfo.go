package models

import (
	"encoding/json"
	"spoty/internal/api"
	"strings"
)

type AlbumFull struct {
	AlbumInfo AlbumInfo `json:"album"`
}

type AlbumInfo struct {
	Artist string `json:"artist"`
	Mbid   string `json:"mbid"`
	Tags   struct {
		Tag []Tag `json:"tag"`
	} `json:"tags"`
	PlayCount string  `json:"playcount"`
	Images    []Image `json:"image"`
	Tracks    struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	Listeners string `json:"listeners"`
	Wiki      Wiki   `json:"wiki"`
}

type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Tag struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type Track struct {
	Streamable struct {
		Fulltrack string `json:"fulltrack"`
		Text      string `json:"#text"`
	}
	Duration   int    `json:"duration"`
	Url        string `json:"url"`
	Name       string `json:"name"`
	Attributes struct {
		Rank int `json:"rank"`
	}
	Artist struct {
		Url  string `json:"url"`
		Name string `json:"name"`
		Mbid string `json:"mbid"`
	} `json:"artist"`
}

func GetAlbumInfo(artistName, albumName string) (AlbumInfo, error) {
	var info AlbumFull

	request1 := strings.ReplaceAll(artistName, " ", "+")
	request2 := strings.ReplaceAll(albumName, " ", "+")
	body, err := api.ParseJson("http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=" + API_KEY + "&artist=" + request1 + "&album=" + request2 + "&format=json")
	err = json.Unmarshal(body, &info)

	return info.AlbumInfo, err
}
