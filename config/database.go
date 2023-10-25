package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

func ConnectDb() (*sql.DB, error) {
	// Konfigurasi koneksi database
	dsn := "root:@tcp(localhost:3306)/aktif_hutan"

	// Membuat koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Sukses terhubung ke database!")
	return db, nil
}
