package homemodel

import "database/sql"

type KegiatanModel struct {
	db *sql.DB
}

func NewKegiatanModel(db *sql.DB) *KegiatanModel {
	return &KegiatanModel{db}
}

type Kegiatan struct {
	KegiatanId     int
	KomunitasNama  string
	Judul_kegiatan string
	Deskripsi      string
	KomunitasId    int
}

func (km *KegiatanModel) GetKegiatanData() ([]Kegiatan, error) {
	query := "SELECT KegiatanId, KomunitasNama, Judul_kegiatan, Deskripsi, KomunitasId FROM kegiatan"

	rows, err := km.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kegiatans []Kegiatan

	for rows.Next() {
		var kegiatan Kegiatan
		err := rows.Scan(&kegiatan.KegiatanId, &kegiatan.KomunitasNama, &kegiatan.Judul_kegiatan, &kegiatan.Deskripsi, &kegiatan.KomunitasId)
		if err != nil {
			return nil, err
		}
		kegiatans = append(kegiatans, kegiatan)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return kegiatans, nil
}
