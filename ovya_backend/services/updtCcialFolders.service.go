package services

import (
	"database/sql"
	"fmt"
	"time"

	"ovya_backend/model"
)

func GetAllFoldersHistory(db *sql.DB) ([]model.FoldersHistory, error) {
	rows, err := db.Query(`
		SELECT id, dossier_id, ccial_id, date_start, date_end
		FROM dossier_historique
		ORDER BY date_start DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historiques []model.FoldersHistory

	for rows.Next() {
		var h model.FoldersHistory
		err := rows.Scan(&h.ID, &h.DossierID, &h.CcialID, &h.DateStart, &h.DateEnd)
		if err != nil {
			return nil, err
		}
		historiques = append(historiques, h)
	}

	return historiques, nil
}

func UpdateDossierCommercial(db *sql.DB, dossierID int, newCcialID int) error {
	_, err := db.Exec(`
		UPDATE dossier_historique
		SET date_end = $1
		WHERE dossier_id = $2 AND date_end IS NULL
	`, time.Now(), dossierID)
	if err != nil {
		return fmt.Errorf("failed to close previous assignment: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO dossier_historique (dossier_id, ccial_id)
		VALUES ($1, $2)
	`, dossierID, newCcialID)
	if err != nil {
		return fmt.Errorf("failed to insert new assignment: %v", err)
	}

	_, err = db.Exec(`
		UPDATE dossier
		SET ccial_id = $1
		WHERE id = $2
	`, newCcialID, dossierID)
	if err != nil {
		return fmt.Errorf("failed to update dossier table: %v", err)
	}

	return nil
}