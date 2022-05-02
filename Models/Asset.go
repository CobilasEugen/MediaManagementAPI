package Models

import (
	"time"
)

type Asset struct {
	Id int `gorm:"not null;unique;primaryKey"`
	Name string `gorm:"not null"`
	Description string `gorm:"not null"`
	Created time.Time `gorm:"not null"`
	Modified time.Time `gorm:"not null"`
	BrodcastStatus string 
	ContentType string 
	Tennant string `gorm:"not null"`
}