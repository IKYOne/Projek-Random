package controller

import (
	"Lab-Hacking/models"
	"Lab-Hacking/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(models.ResponseServer{
				Status:  "Error",
				Message: "Method tidak diizinkan",
			})
			return

		}

		data, err := repository.GetMahasiswa(db)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ResponseServer{
				Status:  "Error",
				Message: "Terjadi Kesalahan",
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func PostContoller(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(models.ResponseServer{
				Status:  "Error",
				Message: "Method tidak diizinkan",
			})
			return
		}

		var mahasiswaBaru models.Mahasiswa
		if err := json.NewDecoder(r.Body).Decode(&mahasiswaBaru); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ResponseServer{
				Status:  "Error",
				Message: "Format json tidak valid",
			})
			return
		}

		data, err := repository.PostMahasiswa(db, &mahasiswaBaru)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ResponseServer{
				Status:  "Error",
				Message: "Terjadi kesalahan",
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)

	}
}
