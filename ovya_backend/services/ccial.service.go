package services

import (
	"database/sql"
	"fmt"
	"ovya_backend/model"
)

func GetCcialByEmail(db *sql.DB, email string) (model.Ccial, error) {
	var ccial model.Ccial

	err := db.QueryRow("SELECT * FROM ccial WHERE email = $1", email).Scan(&ccial.Id, &ccial.Nom, &ccial.Email)

	if err != nil {
		fmt.Println("error")
		return model.Ccial{}, err
	}
	return ccial, nil
}

func GetAllCcial(db *sql.DB) ([]model.Ccial, error) {
	rows, err := db.Query("SELECT * FROM ccial")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ccial []model.Ccial

	for rows.Next() {
		var user model.Ccial
		err := rows.Scan(&user.Id, &user.Nom, &user.Email)
		if err != nil {
			return nil, err
		}
		ccial = append(ccial, user)
	}

	return ccial, rows.Err()
}

func CreateCcial(db *sql.DB, Ccial model.Ccial) error {
	_, err := db.Exec("INSERT INTO ccial (nom, email) VALUES ($1, $2)", Ccial.Nom, Ccial.Email)

	return err
}

func UpdateCcialEMail(db *sql.DB, newEmail string, id int) error {
	_, err := db.Exec("UPDATE ccial SET email = $1 WHERE id = $2", newEmail, id)

	return err
}

func UpdateCcialName(db *sql.DB, newName string, id int) error {
	_, err := db.Exec("UPDATE ccial SET nom = $1 WHERE id = $2", newName, id)

	return err
}

func DeleteCcial(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM ccial WHERE id = $1", id)

	return err
}
