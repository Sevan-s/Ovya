package routes

import (
	"database/sql"
	"net/http"
	"ovya_backend/controllers"
)

func CreateVisiteRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/create", controllers.CreateVisiteHandler(db))
}

func GetAllVisiteRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/all", controllers.GetAllVisiteHandler(db))
}

func GetVitisteByIdRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite", controllers.GetVisiteByIdHandler(db))
}

func DeleteVisiteRoute(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/delete", controllers.DeleteVisiteHandler(db))
}

func UpdateDateEnd(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updatedateend", controllers.UpdateVisiteDateEndByIdHandler(db))
}

func UpdateDateStart(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updatedatestart", controllers.UpdateVisiteDateStartByIdHandler(db))
}

func UpdateAcqId(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updateacqid", controllers.UpdateVisiteAcqByIdHandler(db))
}

func UpdateCcialId(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updateccialid", controllers.UpdateVisiteCcialByIdHandler(db))
}

func UpdateFolderId(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updatefolderid", controllers.UpdateVisiteFolderByIdHandler(db))
}

func UpdateStatus(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/updateStatus", controllers.UpdateVisiteStatusHandler(db))
}

func GetVisiteByQuantity(db *sql.DB, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/visite/number", controllers.HandleGetVisiteStats(db))
}
