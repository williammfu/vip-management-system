package model

import (
	"time"

	"github.com/williammfu/vip-management-system/utils"
	"gorm.io/gorm"
)

type Vip struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	CountryOfOrigin string    `json:"country_of_origin"`
	ETA             time.Time `json:"eta"`
	Photo           string    `json:"photo"`
	Attributes      []string  `json:"attributes"`
}

var db *gorm.DB

func (v Vip) GetGuestInfo() (g Guest) {
	return Guest{
		ID:              int(utils.HashID(v.Name)),
		Name:            v.Name,
		CountryOfOrigin: v.CountryOfOrigin,
		ETA:             v.ETA,
		Photo:           v.Photo,
		Arrived:         false}
}

func (v Vip) GetDescriptions(id int) (d []Description) {
	var td []Description
	for _, atr := range v.Attributes {
		td = append(td, Description{id, atr})
	}
	return td
}

func CreateVip(g Guest, d []Description) Vip {
	var ta []string
	for _, atr := range d {
		ta = append(ta, atr.Attribute)
	}
	return Vip{g.ID, g.Name, g.CountryOfOrigin, g.ETA, g.Photo, ta}
}
