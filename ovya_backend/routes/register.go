package routes

import (
	"database/sql"
	"net/http"
)

func RegisterAllRoutes(db *sql.DB, mux *http.ServeMux) {

	///////////////// Acq routes /////////////////

	RegisterAcqRoutes(db, mux)
	GetAcqRoutes(db, mux)
	GetAcqByNameRoutes(db, mux)
	GetAllAcqRoutes(db, mux)
	DeleteUser(db, mux)
	UpdateAcqEMail(db, mux)
	UpdateAcqName(db, mux)
	UpdateAcqPassword(db, mux)

	///////////////// Ccial routes /////////////////

	GetCcialByEmailRoute(db, mux)
	GetAllCcialFoldersHistory(db, mux)
	CreateCcial(db, mux)
	GetAllCcial(db, mux)
	UpdateCcialName(db, mux)
	UpdateCcialEmail(db, mux)
	DeleteCcial(db, mux)
	UpdateCcialFolderHistory(db, mux)
	///////////////// Folder routes /////////////////

	CreateFolderRoute(db, mux)
	GetFolderRoute(db, mux)
	GetAllFolderRoute(db, mux)
	UpdateFolderRoute(db, mux)
	DeleteFolderRoute(db, mux)
	SearchFolderRoute(db, mux)

	///////////////// Visite routes /////////////////
	CreateVisiteRoute(db, mux)
	GetAllVisiteRoute(db, mux)
	GetVitisteByIdRoute(db, mux)
	DeleteVisiteRoute(db, mux)
	UpdateDateEnd(db, mux)
	UpdateDateStart(db, mux)
	UpdateAcqId(db, mux)
	UpdateCcialId(db, mux)
	UpdateFolderId(db, mux)
	UpdateStatus(db, mux)
	GetVisiteByQuantity(db, mux)
}
