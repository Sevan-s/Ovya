package routes

import (
	"database/sql"
	"net/http"
	"ovya_backend/controllers"
)

func RegisterAcqRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/create", controllers.CreateAcqHandler(db))
}

func GetAcqRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq", controllers.GetUserByEmailHandler(db))
}
func GetAcqByNameRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/name", controllers.GetUserByNameHandler(db))
}

func GetAllAcqRoutes(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/all", controllers.GetAllAcqHandler(db))
}

func DeleteUser(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/delete", controllers.DeleteAcqHandler(db))
}

func UpdateAcqEMail(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/updateemail", controllers.UpdateAcqEmailHandler(db))
}

func UpdateAcqName(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/updateename", controllers.UpdateAcqNameHandler(db))
}

func UpdateAcqPassword(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/acq/updatepassword", controllers.UpdateAcqPasswordHandler(db))
}
