package models

type Platform struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Cost      string `json:"cost"`
}
