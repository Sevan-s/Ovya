package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ovya_backend/middleware"
	"ovya_backend/model"
	"ovya_backend/services"
)

func GetUserByEmailHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		email := req.URL.Query().Get("email")
		if email == "" {
			http.Error(res, "Email is required", http.StatusBadRequest)
			return
		}

		acq, err := services.GetUserByEmail(db, email)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(res, "User not found", http.StatusNotFound)
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(acq)
	}
}

func GetUserByNameHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("nom")
		if name == "" {
			http.Error(res, "Name is required", http.StatusBadRequest)
			return
		}

		acq, err := services.GetUserByName(db, name)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(res, "User not found", http.StatusNotFound)
				return
			}
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(acq)
	}
}

func GetAllAcqHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		acq, err := services.GetAllAcq(db)
		if err != nil {
			http.Error(res, "Error retrieving data", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(acq)
	}
}

func CreateAcqHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var acq model.Acq
		err := json.NewDecoder(req.Body).Decode(&acq)
		if err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err = services.CreateUser(db, acq)
		if err != nil {
			http.Error(res, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(map[string]string{"message": "User created successfully"})
	}

}

func DeleteAcqHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method received:", r.Method)
		id, err := middleware.ExtractQueryId(r)
		if err != nil {
			fmt.Println("ID extraction error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = services.DeleteAcq(db, id)
		if err != nil {
			http.Error(w, "Failed to delete User", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User deleted",
		})
	}
}

func UpdateAcqEmailHandler(db *sql.DB) http.HandlerFunc {
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

		err = services.UpdateAcqEMail(db, body.Email, id)

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

func UpdateAcqNameHandler(db *sql.DB) http.HandlerFunc {
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

		err = services.UpdateAcqName(db, body.Nom, id)

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

func UpdateAcqPasswordHandler(db *sql.DB) http.HandlerFunc {
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
			Password string `json:"password"`
		}

		err = json.NewDecoder(req.Body).Decode(&body)

		if err != nil || body.Password == "" {
			http.Error(res, "Invalid password format", http.StatusBadRequest)
			return
		}

		err = services.UpdateAcqPassword(db, body.Password, id)

		if err != nil {
			http.Error(res, "Failed to update password", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"message": "password updated successfully",
		})
	}
}
