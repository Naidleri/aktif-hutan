package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

func ConnectDb() {
	// Konfigurasi koneksi database
	dsn := "root:@tcp(localhost:3306)/aktif_hutan"

	// Membuat koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Gagal menghubungkan ke database:", err)
		return
	}
	defer db.Close()

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		fmt.Println("Gagal menghubungkan ke database:", err)
		return
	}

	log.Println("Sukses terhubung ke database!")
}
