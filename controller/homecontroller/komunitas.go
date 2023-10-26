package homecontroller

import (
	"html/template"
	"net/http"
)

func Komunitas(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/home/komunitas.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}
