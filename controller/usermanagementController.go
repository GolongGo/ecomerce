package controller

import (
	"encoding/json"
	"net/http"

	// "strconv"
	// "github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	
	"ecomerce/helper"
	"ecomerce/model"
	"ecomerce/config"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// cek metohod yang di kirim (POST)
	if r.Method != http.MethodPost {
		ResponseError(w,http.StatusMethodNotAllowed,"method tidak sesuai")
		return
	}
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseError(w,http.StatusBadRequest, "request body belum sesuai")
		return
	}

	if user.Name == "" || user.Password == "" || len(user.Roles) == 0 {
		ResponseError(w,http.StatusBadRequest, "Username, Password, dan minimal satu Role harus diisi")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError,"Failed to encrypt password")
		return
	}
	user.Password = string(hashedPassword)

	var existingUser model.User
	result := config.DB.Where("email = ?", user.Email).First(&existingUser)

		if result != nil {
		tx := config.DB.Begin()

		for _, role := range user.Roles {
			tx.FirstOrCreate(&model.Role{}, model.Role{Name: role.Name}).Scan(&role)

			tx.Model(&user).Association("Roles").Append(role)
		}
		tx.Create(&user)
		tx.Commit()

		ResponseJson(w,http.StatusCreated,"Success")
	} else{
		ResponseError(w, http.StatusBadRequest, "User Sudah ada")
	}
}
