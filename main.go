package main

import (
	"aktif_hutan/config"
	"aktif_hutan/controller/homecontroller"
	"aktif_hutan/controller/logincontroller"
	"aktif_hutan/controller/profile"
	"net/http"
)

func main() {
	config.ConnectDb()

	http.HandleFunc("/", homecontroller.Home)
	http.HandleFunc("/komunitas", homecontroller.Komunitas)

	http.HandleFunc("/profile", profile.Profile)
	http.HandleFunc("/kegiatan", profile.Kegiatan)

	http.HandleFunc("/register", logincontroller.Registration)
	http.HandleFunc("/login", logincontroller.Login)

	http.ListenAndServe(":3000", nil)

}
