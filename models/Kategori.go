package models

import "net/http"

type Mahasiswa struct {
	ID      int    `json:"id"`
	Nim     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

type ResponseServer struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RecServer struct {
	http.ResponseWriter
	StatusCode int
}

func (rec *RecServer) WriteHeader(code int) {
	rec.ResponseWriter.WriteHeader(code)
	rec.StatusCode = code
}
