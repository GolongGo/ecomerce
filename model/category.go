package model

type Category struct {
	CategoryId int    `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name       string `gorm:"size:255; null ;varchar" json:"name"`
	Type       string `gorm:"size:255; null ;varchar" json:"type"`
}
