package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

var db *sql.DB

func ConnectDb() *sql.DB {
	// Konfigurasi koneksi database
	dsn := "root:@tcp(localhost:3306)/aktif_hutan"

	// Membuat koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Gagal menghubungkan ke database:", err)
		return nil
	}

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		log.Println("Gagal menghubungkan ke database:", err)
		return nil
	}

	log.Println("Sukses terhubung ke database!")
	return db
}
