package model

import (
	"time"
)

type SearchRequest struct {
	Request string
	Search  string
	Page    string
}

type SearchResponse struct {
	Search       []Movie `json:"Search"`
	Total        string  `json:"totalResults"`
	ErrorMessage string  `json:"error_message,omitempty"`
}

type DetailRequest struct {
	Id string
}

type Movie struct {
	Title   string `json:"title"`
	Year    string `json:"year"`
	MovieID string `json:"imdbID"`
	Type    string `json:"type"`
	Poster  string `json:"Poster"`
}

type GetMovieResponse struct {
	Detail Movie `json:"detail"`
}

type Logger struct {
	Request   string
	Response  string
	Timestamp time.Time
}
