package model

import "time"

type Blog struct {
	Id        uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Header    string `gorm:"size:255;not null ;varchar" json:"header"`
	Title     string `gorm:"size:255;not null ;varchar" json:"title"`
	Body      string `gorm:"not null ;longtext" json:"title"`
	Foto      string `gorm:"size:255;not null ;varchar" json:"foto"`
	Category  string `gorm:"size:255;varchar" json:"category"`
	CreatedBy string `gorm:"size:255;varchar" json:"created_by"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UpdatedBy string `gorm:"size:255;varchar" json:"updated_by"`
}
