package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"strconv"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"

	"ecomerce/config"
	"ecomerce/helper"
	"ecomerce/model"
	"ecomerce/structs"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func createUser(w http.ResponseWriter, r *http.Request) {

	// cek metohod yang di kirim (POST)
	if r.Method != http.MethodPost {
		ResponseError(w, http.StatusMethodNotAllowed, "method tidak sesuai")
		return
	}
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "request body belum sesuai")
		return
	}

	if user.Name == "" || user.Password == "" || len(user.Roles) == 0 {
		ResponseError(w, http.StatusBadRequest, "Username, Password, dan minimal satu Role harus diisi")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Failed to encrypt password")
		return
	}
	user.Password = string(hashedPassword)

	var existingUser model.User
	result := config.DB.Where("email = ?", user.Email).First(&existingUser)

	if result != nil {
		tx := config.DB.Begin()

		if len(user.Roles) > 0 {
			for i, role := range user.Roles {
				var existingRole model.Role
				config.DB.FirstOrCreate(&existingRole, model.Role{ID: role.ID})
				user.Roles[i] = existingRole
			}
		}
		tx.Create(&user)
		tx.Commit()

		ResponseJson(w, http.StatusCreated, "Success")
	} else {
		ResponseError(w, http.StatusBadRequest, "User Sudah ada")
	}
}

func editUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "request body belum sesuai")
		return
	}

	var existingUser model.User
	result := config.DB.Preload("Roles").First(&existingUser, userID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	existingUser.Name = user.Name
	existingUser.Email = user.Email

	existingUser.Roles = user.Roles

	config.DB.Save(&existingUser)

		ResponseJson(w, http.StatusOK,"oke")
		return
}


func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId, ok := vars["id"]
	if !ok || UserId == "" {
		ResponseError(w, http.StatusBadRequest, "user_id Wajib diisi")
		return
	}

	var user model.User
	if config.DB.Preload("Roles").First(&user, UserId) == nil {
		ResponseError(w, http.StatusNotFound, "tidak ada data")
		return
	}
	userDataResponse := structs.UserDataResponse{
		ID:        uint(user.UserId),
		Name:      user.Name,
		Email:     user.Email,
		Roles:     user.Roles,
		CreatedAt: user.CreatedAt,
		Status : user.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	ResponseJson(w, http.StatusOK, userDataResponse)
}


func deleteUserbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId, ok := vars["id"]
	if !ok || UserId == "" {
		ResponseError(w, http.StatusBadRequest, "user_id Wajib diisi")
		return
	}

	var user model.User
	if config.DB.First(&user, UserId) == nil {
		ResponseError(w, http.StatusNotFound, "tidak ada data")
		return
	}

	tx := config.DB.Begin()

	tx.Delete(&user)
	tx.Commit()
	ResponseJson(w, http.StatusOK, "success")
}


func getListUser(w http.ResponseWriter, r *http.Request) {

	var paginationParams helper.PaginationParams
	err := schema.NewDecoder().Decode(&paginationParams, r.URL.Query())
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "request yang di kirim tidak sesuai")
		return
	}
	if paginationParams.Page < 1 {
		paginationParams.Page = 1
	}
	if paginationParams.PageSize < 1 {
		paginationParams.PageSize = 10
	}

	offset := (paginationParams.Page - 1) * paginationParams.PageSize

	var users []model.User
	dbQuery := config.DB.Preload("Roles").Offset(offset).Limit(paginationParams.PageSize)

	if paginationParams.Search != "" {
		fields := strings.Split(paginationParams.Search, ",")
		for _, field := range fields {
			dbQuery = dbQuery.Where(fmt.Sprintf("name LIKE ? OR email LIKE ?", "'%"+field+"%'"))
		}
	}

	result := dbQuery.Find(&users)
	var totalRecords int64
	dbQuery.Model(&model.User{}).Count(&totalRecords)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ResponseError(w, http.StatusInternalServerError, "Error")
		return
	}
	var userDataResponses []structs.UserDataResponse
	for _, user := range users {
		roles := getRoles(uint(user.UserId))

		userDataResponse := structs.UserDataResponse{
			ID:        uint(user.UserId),
			Name:      user.Name,
			Email:     user.Email,
			Roles:     roles,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
		}
		userDataResponses = append(userDataResponses, userDataResponse)
		
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(paginationParams.PageSize)))

	response := map[string]interface{}{
		"totalRecords": totalRecords,
		"currentPage":  paginationParams.Page,
		"pageSize":     paginationParams.PageSize,
		"totalPages":   totalPages,
		"users":        userDataResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	ResponseJson(w, http.StatusOK, response)
}

func getRoles(userID uint) []model.Role {
	var user model.User
	if config.DB.Preload("Roles").First(&user, userID) == nil {
		return nil
	}
	return user.Roles
}
