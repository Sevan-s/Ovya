package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ovya_backend/middleware"
	"ovya_backend/model"
	"ovya_backend/services"
	"strconv"
	"time"
)

func CreateVisiteHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var newVisite model.Visite

		err := json.NewDecoder(req.Body).Decode(&newVisite)
		if err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}

		if newVisite.DateStart.IsZero() || newVisite.DateEnd.IsZero() {
			http.Error(res, "Invalid date values", http.StatusBadRequest)
			return
		}

		err = services.CreateVisite(db, newVisite)
		if err != nil {
			http.Error(res, "Failed to create visite", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Visite created successfully",
		})
	}
}

func GetVisiteByIdHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqId := req.URL.Query().Get("id")

		if reqId == "" {
			http.Error(res, "Id is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(reqId)

		if err != nil {
			http.Error(res, "Id is invalid", http.StatusBadRequest)
		}
		visite, err := services.FindVisite(db, id)

		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(res).Encode(map[string]string{
					"error": "No visite found with this ID",
				})
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(visite)
	}
}

func GetAllVisiteHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		visite, err := services.GetVisite(db)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(visite)
	}
}

func DeleteVisiteHandler(db *sql.DB) http.HandlerFunc {
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

		err = services.DeleteVisite(db, id)

		if err != nil {
			http.Error(res, "Failed to delete visite", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "visite deleted successfully",
		})
	}
}

func UpdateVisiteDateEndByIdHandler(db *sql.DB) http.HandlerFunc {
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
			DateEnd time.Time `json:"date_end"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateDateEnd(db, body.DateEnd, id)

		if err != nil {
			http.Error(res, "Failed to update Visite Date", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste date_end updated successfully",
		})
	}
}

func UpdateVisiteDateStartByIdHandler(db *sql.DB) http.HandlerFunc {
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
			DateStart time.Time `json:"date_start"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateDateStart(db, body.DateStart, id)

		if err != nil {
			http.Error(res, "Failed to update Visite date_start", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste date_start updated successfully",
		})
	}
}

func UpdateVisiteAcqByIdHandler(db *sql.DB) http.HandlerFunc {
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
			AcqId int `json:"acq_id"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateAcq_id(db, body.AcqId, id)

		if err != nil {
			http.Error(res, "Failed to update acqId", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste acqId updated successfully",
		})
	}
}

func UpdateVisiteCcialByIdHandler(db *sql.DB) http.HandlerFunc {
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

		err = services.UpdateCcial_id(db, body.CcialId, id)

		if err != nil {
			http.Error(res, "Failed to update ccialId", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste ccialId updated successfully",
		})
	}
}

func UpdateVisiteFolderByIdHandler(db *sql.DB) http.HandlerFunc {
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
			FolderId int `json:"dossier_id"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateFolder_id(db, body.FolderId, id)

		if err != nil {
			http.Error(res, "Failed to update folder", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste folder updated successfully",
		})
	}
}

func UpdateVisiteStatusHandler(db *sql.DB) http.HandlerFunc {
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
			Status bool `json:"canceled"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(res, "Invalid body", http.StatusBadRequest)
			return
		}

		err = services.UpdateStatus(db, body.Status, id)

		if err != nil {
			http.Error(res, "Failed to update status", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Viste status updated successfully",
		})
	}
}

func HandleGetVisiteStats(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		stats, err := services.GetVisiteIfFivePerDay(db)
		if err != nil {
			http.Error(w, "failed to retrieve visite stats", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats)
	}
}
