package Models

import "time"

type AssertFiles struct {
	Id       int       `gorm:"primaryKey;not null"`
	Asset_id int       `gorm:"not null"`
	Created  time.Time `gorm:"not null"`
	Modified time.Time `gorm:"not null"`
	Name string `gorm:"size:50;not null"`
	FileType string `gorm:"size:50;not null"`
	FileSize int64 `gorm:"not null"`
	Profile string `gorm:"size:50"`
	Location string `gorm:"size:255"`
	Tenant string `gorm:"size:50;not null"`
}