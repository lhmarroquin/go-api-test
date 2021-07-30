package main

import "encoding/json"

type Tvmaze []TvmazeElement

func UnmarshalTvmaze(data []byte) (Tvmaze, error) {
	var r Tvmaze
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tvmaze) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvmazeElement struct {
	Score float64 `json:"score"`
	Show  Show    `json:"show"`
}

type Show struct {
	ID             int64       `json:"id"`
	URL            string      `json:"url"`
	Name           string      `json:"name"`
	Type           Type        `json:"type"`
	Language       string      `json:"language"`
	Genres         []string    `json:"genres"`
	Status         Status      `json:"status"`
	Runtime        *int64      `json:"runtime"`
	AverageRuntime *int64      `json:"averageRuntime"`
	Premiered      *string     `json:"premiered"`
	OfficialSite   *string     `json:"officialSite"`
	Schedule       Schedule    `json:"schedule"`
	Rating         Rating      `json:"rating"`
	Weight         int64       `json:"weight"`
	Network        *Network    `json:"network"`
	WebChannel     *Network    `json:"webChannel"`
	DVDCountry     interface{} `json:"dvdCountry"`
	Externals      Externals   `json:"externals"`
	Image          *Image      `json:"image"`
	Summary        *string     `json:"summary"`
	Updated        int64       `json:"updated"`
	Links          Links       `json:"_links"`
}

type Externals struct {
	Tvrage  *int64  `json:"tvrage"`
	Thetvdb *int64  `json:"thetvdb"`
	Imdb    *string `json:"imdb"`
}

type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type Links struct {
	Self            Previousepisode  `json:"self"`
	Previousepisode *Previousepisode `json:"previousepisode,omitempty"`
}

type Previousepisode struct {
	Href string `json:"href"`
}

type Network struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Country *Country `json:"country"`
}

type Country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Timezone string `json:"timezone"`
}

type Rating struct {
	Average *float64 `json:"average"`
}

type Schedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}

type Status string

const (
	Ended         Status = "Ended"
	InDevelopment Status = "In Development"
	Running       Status = "Running"
)

type Type string

const (
	Scripted Type = "Scripted"
)
