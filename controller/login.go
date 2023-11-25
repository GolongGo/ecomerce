package controller

import (
	"ecomerce/config"
	"ecomerce/model"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func hashPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashedPassword
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// err := json.NewDecoder(r.Body).Decode(&user)

	user.Name = r.FormValue("Name")
	var userName = r.FormValue("Name")
	user.Password = r.FormValue("Password")

	fmt.Println(user)
	fmt.Println(userName)

	// if err != nil {
	// 	ResponseError(w, http.StatusInternalServerError, "Failed to encrypt password")
	// 	return
	// }

	result := config.DB.Where("name =?", user.Name)

	fmt.Println(result)

	if result != nil {
		tx := config.DB.Begin()

		tx.Commit()
		ResponseJson(w, http.StatusOK, "succes")
		return
	} else {
		ResponseError(w, http.StatusNotFound, "data tidak di temukan")
		return
	}
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Set authenticated menjadi false
	session.Values["authenticated"] = false
	session.Save(r, w)

	// Redirect ke halaman login setelah logout
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Set authenticated menjadi false
	session.Values["authenticated"] = false
	session.Save(r, w)
}
