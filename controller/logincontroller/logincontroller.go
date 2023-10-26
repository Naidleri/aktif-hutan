package logincontroller

import (
	"aktif_hutan/config"
	"aktif_hutan/model/loginmodel"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/login/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		db := config.ConnectDb()
		userModel := loginmodel.NewUserModel(db)

		authenticated, err := userModel.AuthenticateUser(email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if authenticated {
			// Redirect to the user's dashboard or another page on successful login
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			errorMsg := "Login gagal, silahkan coba lagi."
			html := "<html><body><p>" + errorMsg + "</p></body></html>"
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(html))
		}
	}
}
