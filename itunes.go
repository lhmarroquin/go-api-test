package main

import "encoding/json"

func UnmarshalItunes(data []byte) (Itunes, error) {
	var r Itunes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Itunes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Itunes struct {
	ResultCount int64    `json:"resultCount"`
	Results     []Result `json:"results"`
}

type Result struct {
	WrapperType             *string  `json:"wrapperType"`
	Kind                    *string  `json:"kind,omitempty"`
	TrackID                 *int64   `json:"trackId,omitempty"`
	ArtistName              *string  `json:"artistName"`
	TrackName               *string  `json:"trackName,omitempty"`
	TrackCensoredName       *string  `json:"trackCensoredName,omitempty"`
	TrackViewURL            *string  `json:"trackViewUrl,omitempty"`
	PreviewURL              string   `json:"previewUrl"`
	ArtworkUrl30            *string  `json:"artworkUrl30,omitempty"`
	ArtworkUrl60            string   `json:"artworkUrl60"`
	ArtworkUrl100           string   `json:"artworkUrl100"`
	CollectionPrice         *float64 `json:"collectionPrice,omitempty"`
	TrackPrice              *float64 `json:"trackPrice,omitempty"`
	TrackRentalPrice        *float64 `json:"trackRentalPrice,omitempty"`
	CollectionHDPrice       *float64 `json:"collectionHdPrice,omitempty"`
	TrackHDPrice            *float64 `json:"trackHdPrice,omitempty"`
	TrackHDRentalPrice      *float64 `json:"trackHdRentalPrice,omitempty"`
	ReleaseDate             string   `json:"releaseDate"`
	CollectionExplicitness  *string  `json:"collectionExplicitness"`
	TrackExplicitness       *string  `json:"trackExplicitness,omitempty"`
	TrackTimeMillis         *int64   `json:"trackTimeMillis,omitempty"`
	Country                 *string  `json:"country"`
	Currency                *string  `json:"currency"`
	PrimaryGenreName        *string  `json:"primaryGenreName"`
	ContentAdvisoryRating   *string  `json:"contentAdvisoryRating,omitempty"`
	ShortDescription        *string  `json:"shortDescription,omitempty"`
	LongDescription         *string  `json:"longDescription,omitempty"`
	ArtistID                *int64   `json:"artistId,omitempty"`
	CollectionID            *int64   `json:"collectionId,omitempty"`
	CollectionName          *string  `json:"collectionName,omitempty"`
	CollectionCensoredName  *string  `json:"collectionCensoredName,omitempty"`
	ArtistViewURL           *string  `json:"artistViewUrl,omitempty"`
	CollectionViewURL       *string  `json:"collectionViewUrl,omitempty"`
	TrackCount              *int64   `json:"trackCount,omitempty"`
	Copyright               *string  `json:"copyright,omitempty"`
	Description             *string  `json:"description,omitempty"`
	DiscCount               *int64   `json:"discCount,omitempty"`
	DiscNumber              *int64   `json:"discNumber,omitempty"`
	TrackNumber             *int64   `json:"trackNumber,omitempty"`
	IsStreamable            *bool    `json:"isStreamable,omitempty"`
	CollectionArtistID      *int64   `json:"collectionArtistId,omitempty"`
	CollectionArtistName    *string  `json:"collectionArtistName,omitempty"`
	CollectionArtistViewURL *string  `json:"collectionArtistViewUrl,omitempty"`
}
