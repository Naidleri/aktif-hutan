package loginmodel

import (
	"database/sql"
	_ "errors"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db}
}

func (um *UserModel) AuthenticateUser(email string, password string) (bool, error) {
	query := "SELECT COUNT(*) FROM komunitas WHERE Email = ? AND Password = ?"
	var count int
	err := um.db.QueryRow(query, email, password).Scan(&count)

	if err != nil {
		return false, err
	}

	// Jika count > 0, berarti ada data yang sesuai
	return count > 0, nil
}
