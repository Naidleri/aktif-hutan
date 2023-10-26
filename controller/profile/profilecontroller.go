package profile

import (
	"html/template"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/profile/profile.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
