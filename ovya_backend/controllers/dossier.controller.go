package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ovya_backend/middleware"
	"ovya_backend/model"
	"ovya_backend/services"
	"strconv"
)

func CreateFolderHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var newFolder model.Folder

		err := json.NewDecoder(req.Body).Decode(&newFolder)
		if err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}

		folderId, err := services.CreateFolder(db, newFolder)
		if err != nil {
			http.Error(res, "Failed to create folder", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]interface{}{
			"message":   "Folder created successfully",
			"folder_id": folderId,
		})
	}
}

func GetFolderByIdHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id, err := middleware.ExtractQueryId(req)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		folder, err := services.GetFolderById(db, id)

		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(res).Encode(map[string]string{
					"error": "No folder found with this ID",
				})
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(folder)
	}
}

func GetAllFolderHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		folder, err := services.GetAllFolders(db)

		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(res).Encode(map[string]string{
					"error": "No folder found with this ID",
				})
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(folder)
	}
}

func UpdateFolderByIdHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := middleware.ExtractQueryId(req)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		var body struct {
			CcialId int `json:"ccial_id"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateFolderCcial(db, id, body.CcialId)

		if err != nil {
			http.Error(res, "Failed to update folder", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Folder updated successfully",
		})
	}
}

func DeleteFolderHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqId := req.URL.Query().Get("id")

		if reqId == "" {
			http.Error(res, "Missing Id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(reqId)

		if err != nil {
			http.Error(res, "Invalid Id", http.StatusBadRequest)
			return
		}

		err = services.DeleteFolder(db, id)

		if err != nil {
			http.Error(res, "Failed to delete Folder", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Folder deleted successfully",
		})
	}
}

func SearchFolderHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query().Get("query")
		if query == "" {
			http.Error(res, "Missing query", http.StatusBadRequest)
			return
		}

		folders, err := services.SearchFolders(db, query)
		if err != nil {
			http.Error(res, "Search failed", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(folders)
	}
}
