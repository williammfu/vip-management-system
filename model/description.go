package model

type Description struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Attribute string `json:"attribute" gorm:"primaryKey"`
}
