package logincontroller

import (
	"aktif_hutan/config"
	"aktif_hutan/model/loginmodel"
	"html/template"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/login/register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		email := r.FormValue("email")
		password := r.FormValue("password")
		deskripsi := r.FormValue("deskripsi")
		kabupaten := r.FormValue("kabupaten")
		provinsi := r.FormValue("provinsi")

		db := config.ConnectDb() // Mencoba menghubungkan ke database

		if db == nil {
			http.Error(w, "Gagal menghubungkan ke database", http.StatusInternalServerError)
			return
		}

		model := loginmodel.NewRegistrationModel(db)
		success := model.RegisterUser(nama, email, password, deskripsi, kabupaten, provinsi)

		if success {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "Gagal membuat akun", http.StatusInternalServerError)
		}
	}
}
