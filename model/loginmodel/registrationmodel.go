package loginmodel

import (
	"database/sql"
)

type RegistrationModel struct {
	db *sql.DB
}

func NewRegistrationModel(db *sql.DB) *RegistrationModel {
	return &RegistrationModel{db}
}

func (rm *RegistrationModel) RegisterUser(nama string, email string, password string, deskripsi string, kabupaten string, provinsi string) bool {
	// Mulai transaksi
	result, err := rm.db.Exec("INSERT INTO komunitas (Nama, Email, Password, Deskripsi, Kabupaten, Provinsi) VALUES (?, ?, ?, ?, ?, ?)",
		nama, email, password, deskripsi, kabupaten, provinsi)

	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}
