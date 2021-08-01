package model

import (
	"time"
)

type Guest struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	CountryOfOrigin string    `json:"country_of_origin"`
	ETA             time.Time `json:"eta"`
	Photo           string    `json:"photo"`
	Arrived         bool      `json:"arrived"`
}

func (g *Guest) AssignGuest(other Guest) {
	g.ID = other.ID
	g.Name = other.Name
	g.CountryOfOrigin = other.CountryOfOrigin
	g.ETA = other.ETA
	g.Photo = other.Photo
	g.Arrived = other.Arrived
}
