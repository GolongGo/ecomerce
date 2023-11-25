package model

import "time"

type Products struct {
	ProductId   int    `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName string `gorm:"size:255;not null;" json:"product_name"`
	Harga       int    `gorm:"size:255;not null;" json:"harga"`
	Lokasi      string `gorm:"size:255;not null;" json:"lokasi"`
	Date        time.Time
	Foto        string `gorm:"size:255;not null;" json:"foto"`
	Contact     string `gorm:"size:255;not null;" json:"contact"`
	// Category    []CategoryItem `gorm:"many2many:category_item;" json:"category"`
	Detail     string `gorm:"null;longtext" json:"detail"`
	IfoPenting string `gorm:" null;longtext" json:"info_penting"`
	Stok       int    `gorm:"size:255;not null;" json:"name"`
	Terjual    int    `json:"terjual"`
	Rating     int    `json:"rating"`
}
