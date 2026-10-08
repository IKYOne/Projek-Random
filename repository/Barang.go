package repository

import (
	"Lab-Hacking/models"
	"database/sql"
)

func GetMahasiswa(db *sql.DB) (*models.ResponseServer, error) {
	query := "SELECT id, nim, nama, jurusan FROM mahasiswa"
	rews, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rews.Close()

	daftarMahasiswa := make([]*models.Mahasiswa, 0)

	for rews.Next() {
		mahasiswa := &models.Mahasiswa{}
		if err := rews.Scan(&mahasiswa.ID, &mahasiswa.Nim, &mahasiswa.Nama, &mahasiswa.Jurusan); err != nil {
			return nil, err
		}

		daftarMahasiswa = append(daftarMahasiswa, mahasiswa)

	}

	if err := rews.Err(); err != nil {
		return nil, err
	}

	return &models.ResponseServer{
		Status:  "Success",
		Message: "Data Mahasiswa",
		Data:    daftarMahasiswa,
	}, nil

}

func GetMahasiswaById(db *sql.DB, id int) (*models.ResponseServer, error) {
	query := "SELECT id,nim,nama,jurusan FROM mahasiswa WHERE id = $1"

	rows, err := db.QueryRow()
}

func PostMahasiswa(db *sql.DB, data *models.Mahasiswa) (*models.ResponseServer, error) {
	query := "INSERT INTO mahasiswa (nim,nama,jurusan) VALUES ($1,$2,$3)"
	_, err := db.Exec(query, data.Nim, data.Nama, data.Jurusan)
	if err != nil {
		return nil, err
	}

	return &models.ResponseServer{
		Status:  "Success",
		Message: "Data Mahasiswa Berhasi di tambahkan ke database",
	}, nil

}
