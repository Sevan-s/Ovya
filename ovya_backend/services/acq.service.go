package services

import (
	"database/sql"
	"fmt"
	"ovya_backend/model"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(db *sql.DB, email string) (model.Acq, error) {
	var acq model.Acq

	err := db.QueryRow("SELECT * FROM acq WHERE email = $1", email).Scan(&acq.ID, &acq.Nom, &acq.Email, &acq.Password)

	if err != nil {
		fmt.Println("error")
		return model.Acq{}, err
	}
	return acq, err
}

func GetUserByName(db *sql.DB, name string) (model.Acq, error) {
	var acq model.Acq

	err := db.QueryRow("SELECT * FROM acq WHERE nom = $1", name).Scan(&acq.ID, &acq.Nom, &acq.Email, &acq.Password)

	if err != nil {
		fmt.Println("error")
		return model.Acq{}, err
	}
	return acq, err
}

func GetAllAcq(db *sql.DB) ([]model.Acq, error) {
	rows, err := db.Query("SELECT * FROM acq")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acq []model.Acq

	for rows.Next() {
		var user model.Acq
		err := rows.Scan(&user.ID, &user.Nom, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		acq = append(acq, user)
	}

	return acq, rows.Err()
}

func CreateUser(db *sql.DB, acq model.Acq) (sql.Result, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("INSERT INTO acq (nom, email, password) VALUES ($1, $2, $3)", acq.Nom, acq.Email, hashedPassword)

	return result, err
}

func UpdateAcqEMail(db *sql.DB, newEmail string, id int) error {
	_, err := db.Exec("UPDATE acq SET email = $1 WHERE id = $2", newEmail, id)

	return err
}

func UpdateAcqPassword(db *sql.DB, newPassword string, id int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE acq SET password = $1 WHERE id = $2", hashedPassword, id)

	return err
}

func UpdateAcqName(db *sql.DB, newName string, id int) error {
	_, err := db.Exec("UPDATE acq SET nom = $1 WHERE id = $2", newName, id)

	return err
}

func DeleteAcq(db *sql.DB, id int) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM visite WHERE acq_id = $1", id).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("cet acquéreur est lié à %d visite(s)", count)
	}
	_, err = db.Exec("DELETE FROM acq WHERE id = $1", id)
	return err
}
