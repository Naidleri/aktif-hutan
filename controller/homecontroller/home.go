package homecontroller

import (
	"aktif_hutan/config"
	"aktif_hutan/model/homemodel"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := config.ConnectDb()
		kegiatanModel := homemodel.NewKegiatanModel(db)

		// Mengambil data kegiatan dari model
		kegiatans, err := kegiatanModel.GetKegiatanData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Menetapkan data Kegiatans ke tampilan
		data := struct {
			Kegiatans []homemodel.Kegiatan // Sesuaikan dengan struktur model Anda
		}{
			Kegiatans: kegiatans,
		}

		// Eksekusi template dengan data
		temp, err := template.ParseFiles("view/home/homepagenodb.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
	}
}
