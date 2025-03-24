package services

import (
	"database/sql"
	"encoding/json"
	"ovya_backend/model"
	"time"
)

func CreateVisite(db *sql.DB, Visite model.Visite) error {
	_, err := db.Exec("INSERT INTO visite (date_start, date_end, acq_id, ccial_id, dossier_id, canceled) VALUES ($1, $2, $3, $4, $5, $6)", Visite.DateStart, Visite.DateEnd, Visite.AcqId, Visite.CcialId, Visite.DossierId, Visite.Canceled)

	return err
}

func GetVisite(db *sql.DB) ([]model.Visite, error) {
	rows, err := db.Query("SELECT * FROM visite")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visites []model.Visite

	for rows.Next() {
		var v model.Visite
		if err := rows.Scan(&v.Id, &v.DateStart, &v.DateEnd, &v.AcqId, &v.CcialId, &v.DossierId, &v.Canceled); err != nil {
			return nil, err
		}
		visites = append(visites, v)
	}

	return visites, rows.Err()
}

func FindVisite(db *sql.DB, id int) (model.Visite, error) {
	var visite model.Visite

	err := db.QueryRow("SELECT * FROM visite WHERE id = $1", id).
		Scan(&visite.Id, &visite.DateStart, &visite.DateEnd, &visite.AcqId, &visite.CcialId, &visite.DossierId, &visite.Canceled)

	if err != nil {
		return model.Visite{}, err
	}
	return visite, nil
}

func UpdateDateEnd(db *sql.DB, dateEnd time.Time, id int) error {
	_, err := db.Exec("UPDATE visite SET date_end = $1 Where id = $2", dateEnd, id)
	return err
}
func UpdateDateStart(db *sql.DB, dateStart time.Time, id int) error {
	_, err := db.Exec("UPDATE visite SET date_start = $1 Where id = $2", dateStart, id)
	return err
}

func UpdateAcq_id(db *sql.DB, acq_id int, id int) error {
	_, err := db.Exec("UPDATE visite SET acq_id = $1 Where id = $2", acq_id, id)
	return err
}

func UpdateCcial_id(db *sql.DB, ccial_id int, id int) error {
	_, err := db.Exec("UPDATE visite SET ccial_id = $1 Where id = $2", ccial_id, id)
	return err
}

func UpdateFolder_id(db *sql.DB, dossier_id int, id int) error {
	_, err := db.Exec("UPDATE visite SET dossier_id = $1 Where id = $2", dossier_id, id)
	return err
}

func UpdateStatus(db *sql.DB, status bool, id int) error {
	_, err := db.Exec("UPDATE visite SET canceled = $1 Where id = $2", status, id)
	return err
}

func DeleteVisite(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM visite WHERE id = $1", id)
	return err
}

func GetVisiteIfFivePerDay(db *sql.DB) ([]model.VisiteGroup, error) {
	query := `
		SELECT
			DATE(date_start) AS date,
			ccial_id,
			COUNT(*) AS visite_nb,
			json_agg(dossier_id) AS dossier_ids
		FROM visite
		WHERE canceled = false
		GROUP BY DATE(date_start), ccial_id
		HAVING COUNT(*) > 5
		ORDER BY DATE(date_start)
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.VisiteGroup

	for rows.Next() {
		var group model.VisiteGroup
		var dossierIds []byte

		if err := rows.Scan(&group.Date, &group.CcialID, &group.VisiteNb, &dossierIds); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(dossierIds, &group.DossierIds); err != nil {
			return nil, err
		}

		result = append(result, group)
	}

	return result, nil
}
