package main

import (
	"log"
	"net/http"

	"csv-importer/internal/config"
	"csv-importer/internal/database"
	"csv-importer/internal/handler"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("/upload", handler.UploadHandler(db))
	mux.HandleFunc("/tasks", handler.GetTasksHandler(db))
	mux.HandleFunc("/delete-all", handler.DeleteAllHandler(db))

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
