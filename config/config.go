package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"ecomerce/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/ecommerce"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRoles{}, &model.Blog{}, &model.Products{}, &model.CategoryItem{}, &model.Category{})
	DB = db
}
