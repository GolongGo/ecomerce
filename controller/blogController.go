package controller

import (
	"ecomerce/config"
	"ecomerce/model"
	"net/http"
)

func multiple(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		ResponseError(w, http.StatusMethodNotAllowed, "Method tidak Sesuai")
		return
	}
	var blog model.Blog
	var action = r.FormValue("action")
	var id = r.FormValue("Id")
	blog.Header = r.FormValue("Header")
	blog.Title = r.FormValue("Title")
	blog.Body = r.FormValue("Body")
	blog.Foto = r.FormValue("Foto")
	blog.Category = r.FormValue("Category")

	if &blog == nil {
		ResponseError(w, http.StatusBadRequest, "request body belum sesuai")
		return
	}

	var existing model.Blog
	switch action {
	case "create":
		tx := config.DB.Begin()
		tx.Create(&blog)
		tx.Commit()
		ResponseJson(w, http.StatusAccepted, "succes")
	case "update":

		rest := config.DB.Where("id=?", id).First(&existing)

		if rest.RowsAffected == 0 {
			ResponseError(w, http.StatusNoContent, "Data Tidak Ditemukan")
			return
		} else {
			tx := config.DB.Begin()
			tx.Model(&existing).Updates(blog)
			tx.Commit()
			ResponseJson(w, http.StatusOK, "succes")
		}
	case "delete":
		del := config.DB.First(&existing, id)

		if del.RowsAffected == 0 {
			ResponseError(w, http.StatusNotFound, "tidak ada data")
			return
		} else {
			tx := config.DB.Begin()
			tx.Delete(&existing)
			tx.Commit()
			ResponseJson(w, http.StatusOK, "Data Deleted!")
		}

	default:
		ResponseError(w, http.StatusBadRequest, "Action tidak Sesuai")
	}
}
