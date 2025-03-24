package routes

import (
	"database/sql"
	"net/http"
	"ovya_backend/controllers"
)

func GetCcialByEmailRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial", controllers.GetCcialByEmailHandler(db))
}

func GetAllCcial(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/all", controllers.GetAllCcialHandler(db))
}

func GetAllCcialFoldersHistory(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/history", controllers.GetAllCcialFoldersHistorysHandler(db))
}

func CreateCcial(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/create", controllers.CreateCcialHandler(db))
}

func UpdateCcialName(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/updatename", controllers.UpdateCcialNameHandler(db))
}

func UpdateCcialEmail(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/updateemail", controllers.UpdateCcialEmailHandler(db))
}
func UpdateCcialFolderHistory(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/updateCcialIdHistory", controllers.UpdateCcialIdHandler(db))
}

func DeleteCcial(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ccial/delete", controllers.DeleteCcialHandler(db))
}
