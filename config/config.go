package config

import(
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"ecomerce/model"

)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/ecommerce?parseTime=true"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Role{})
	DB = db
}
