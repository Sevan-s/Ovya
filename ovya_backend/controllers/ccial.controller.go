package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ovya_backend/middleware"
	"ovya_backend/model"
	"ovya_backend/services"
	"fmt"
)

func GetCcialByEmailHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		email := req.URL.Query().Get("email")

		if email == "" {
			http.Error(res, "Email is required", http.StatusBadRequest)
			return
		}

		ccial, err := services.GetCcialByEmail(db, email)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(res, "User not found", http.StatusNotFound)
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(ccial)

	}
}

func GetAllCcialHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		acq, err := services.GetAllCcial(db)
		if err != nil {
			http.Error(res, "Error retrieving data", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(acq)
	}
}

func GetAllCcialFoldersHistorysHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		historiques, err := services.GetAllFoldersHistory(db)
		if err != nil {
			http.Error(res, "Error", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(historiques)
	}
}

func CreateCcialHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var newCcial model.Ccial

		err := json.NewDecoder(req.Body).Decode(&newCcial)

		if err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = services.CreateCcial(db, newCcial)

		if err != nil {
			http.Error(res, "failed to create ccial", http.StatusInternalServerError)
			return
		}
		res.Header().Set("content-type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{"message": "User created successfully"})
	}
}

func UpdateCcialEmailHandler(db *sql.DB) http.HandlerFunc {
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
			Email string `json:"email"`
		}
		err = json.NewDecoder(req.Body).Decode(&body)
		if err != nil || body.Email == "" {
			http.Error(res, "Invalid email format", http.StatusBadRequest)
			return
		}
		err = services.UpdateCcialEMail(db, body.Email, id)
		if err != nil {
			http.Error(res, "Failed to update email", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Email updated successfully",
		})
	}
}

func UpdateCcialNameHandler(db *sql.DB) http.HandlerFunc {
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
			Nom string `json:"nom"`
		}

		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil || body.Nom == "" {
			http.Error(res, "Invalid name format", http.StatusBadRequest)
			return
		}

		err = services.UpdateCcialName(db, body.Nom, id)

		if err != nil {
			http.Error(res, "Failed to update name", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"message": "name updated successfully",
		})
	}
}

func UpdateCcialIdHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			DossierID int `json:"dossier_id"`
			CcialID   int `json:"ccial_id"`
		}

		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			http.Error(res, "Requête invalide", http.StatusBadRequest)
			return
		}

		if err := services.UpdateDossierCommercial(db, data.DossierID, data.CcialID); err != nil {
			http.Error(res, fmt.Sprintf("Erreur : %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(res).Encode(map[string]string{
			"message": "Affectation mise à jour avec historique",
		})
	}
}

func DeleteCcialHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id, err := middleware.ExtractQueryId(req)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		err = services.DeleteCcial(db, id)

		if err != nil {
			http.Error(res, "Failed to delete User", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "User deleted successfully",
		})
	}
}
