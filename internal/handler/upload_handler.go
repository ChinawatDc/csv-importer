package handler

import (
	"csv-importer/internal/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

func UploadHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "file missing", 400)
			return
		}
		defer file.Close()

		inserted, err := service.ImportCSV(file, db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"inserted": inserted,
		})
	}
}

func GetTasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tasks, err := service.GetAllTasks(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func DeleteAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		affected, err := service.DeleteAll(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"deleted": affected,
		})
	}
}
