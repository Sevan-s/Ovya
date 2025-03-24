package routes

import (
	"database/sql"
	"net/http"
	"ovya_backend/controllers"
)

func CreateFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder/create", controllers.CreateFolderHandler(db))
}

func GetFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder", controllers.GetFolderByIdHandler(db))
}

func GetAllFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder/all", controllers.GetAllFolderHandler(db))
}

func UpdateFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder/update", controllers.UpdateFolderByIdHandler(db))
}

func DeleteFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder/delete", controllers.DeleteFolderHandler(db))
}

func SearchFolderRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/folder/search", controllers.SearchFolderHandler(db))
}
