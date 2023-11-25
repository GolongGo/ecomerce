package controller

import (
	"ecomerce/config"
	"ecomerce/model"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func login(w http.ResponseWriter, r *http.Request) {

	var email = r.FormValue("Email")
	var pass = r.FormValue("Password")

	var user model.User
	result := config.DB.Where("email =?", email).First(&user)

	if result.Error != nil {
		ResponseError(w, http.StatusInternalServerError, "Data not Found")
		return
	}

	if result.RowsAffected == 0 {
		ResponseError(w, http.StatusNoContent, "User not found")
		return
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(pass))

	if err != nil {
		ResponseError(w, http.StatusForbidden, "Incorrect password")
		return
	}

	session, _ := store.Get(r, "id")
	session.Values["authentecated"] = true
	session.Values["name"] = user.Name
	session.Values["roles"] = user.Roles
	session.Save(r, w)

	ResponseJson(w, http.StatusOK, "login success!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// Hapus nilai "authenticated" dari sesi
	delete(session.Values, "authenticated")

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	ResponseJson(w, http.StatusOK, "logout success!")
}
