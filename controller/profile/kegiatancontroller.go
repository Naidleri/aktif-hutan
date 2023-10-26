package profile

import (
	"html/template"
	"net/http"
)

func Kegiatan(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/profile/kegiatan.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
