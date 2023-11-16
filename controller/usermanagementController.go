package controller

import (
	"encoding/json"
	"net/http"

	// "strconv"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"ecomerce/config"
	"ecomerce/helper"
	"ecomerce/model"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func createUser(w http.ResponseWriter, r *http.Request) {

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

func editUser(w http.ResponseWriter, r *http.Request) {
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
	var existing model.User
	result := config.DB.Where("id_user = ?", existing.UserId).First(&existing)
	 if result == nil{

		tx := config.DB.Begin()

		tx.Model(&existing).Updates(user)

	for _, role := range user.Roles {
		tx.FirstOrCreate(&model.Role{}, model.Role{Name: role.Name}).Scan(&role)
		tx.Model(&existing).Association("Roles").Append(role)
	}

	tx.Commit()
	ResponseJson(w,http.StatusOK,"succes")
	return
	 } else{
		ResponseError(w,http.StatusNotFound,"data tidak di temukan")
		return
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId, ok := vars["user_id"]
	if !ok || UserId == "" {
		ResponseError(w,http.StatusBadRequest,"user_id Wajib diisi")
		return
	}

	var user model.User
	if config.DB.Preload("Roles").First(&user, UserId) != nil {
		ResponseError(w,http.StatusNotFound,"tidak ada data")
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError,"error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ResponseJson(w,http.StatusOK,"succes")
	w.Write(userJSON)
}

func deleteUserbyId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	UserId, ok := vars["user_id"]
	if !ok || UserId == "" {
		ResponseError(w,http.StatusBadRequest,"user_id Wajib diisi")
		return
	}

	var user model.User
	if config.DB.First(&user, UserId) != nil {
		ResponseError(w,http.StatusNotFound,"tidak ada data")
		return
	}

	tx := config.DB.Begin()

	tx.Delete(&user)
	tx.Commit()
	ResponseJson(w,http.StatusOK,"succes")
}
