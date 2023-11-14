package model

import (
	"time"
)

type User struct{

	UserId 	uint32  `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name 	string `gorm:"size:255;not null;unique" json:"nickname"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Status bool`json:"status"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
}
