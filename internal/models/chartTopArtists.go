package models

import (
	"encoding/json"
	"spoty/internal/api"
	"os"
)

var API_KEY = os.Getenv("LASTFM_API_KEY")

type ChartTopArtists struct {
	ArtistsF ArtistsFull `json:"artists"`
}

type ArtistsFull struct {
	Artists    []*Artist `json:"artist"`
	Attributes struct {
		Page       string `json:"page"`
		PerPage    string `json:"perPage"`
		TotalPages string `json:"totalPages"`
		Total      string `json:"total"`
	} `json:"@attr"`
}

type Artist struct {
	Name       string  `json:"name"`
	PlayCount  string  `json:"playcount"`
	Listeners  string  `json:"listeners"`
	Mbid       string  `json:"mbid"`
	Url        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Images     []Image `json:"image"`

	Thumb string
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

func GetChartTopArtists() ([]*Artist, error) {
	var artists ChartTopArtists
	body, err := api.ParseJson("http://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&api_key=" + API_KEY + "&format=json")
	err = json.Unmarshal(body, &artists)

	return artists.ArtistsF.Artists, err
}

func TopChartArtistsList(len int) ([]*Artist, error) {
	arr, err := GetChartTopArtists()
	arr = arr[:len]

	for _, artist := range arr {
		albums, err := GetArtistsTopAlbums(artist.Name)
		if err != nil {
			return nil, err
		}

		artist.Thumb = albums.Albums[0].Images[3].Text
	}

	return arr, err
}
