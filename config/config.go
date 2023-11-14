package config

import(
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"ecomerce/model"

)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/ecomerce"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRoles{})
	DB = db
}
