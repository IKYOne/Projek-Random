package main

import (
	"Lab-Hacking/controller"
	"Lab-Hacking/database"
	"Lab-Hacking/service/middleware"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		slog.Error("gagal terkoneksi database")
		return
	}
	defer db.Close()

	fmt.Println("Database Berhasil terhubung")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.GetController(db))
	mux.HandleFunc("POST /create", controller.PostContoller(db))

	slog.Info("Server berjalan di http://localhost:8080...")
	if err := http.ListenAndServe(":8080", middleware.Logging(mux)); err != nil {
		panic(err)
	}

}
