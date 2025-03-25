package services

import (
	"database/sql"
	"fmt"
	"ovya_backend/model"
	"time"
)

func CreateFolder(db *sql.DB, folder model.Folder) (int, error) {
	var folderId int

	err := db.QueryRow(`
		INSERT INTO dossier (date_insert, ccial_id)
		VALUES ($1, $2)
		RETURNING id
	`, time.Now(), folder.Ccial_Id).Scan(&folderId)

	if err != nil {
		return 0, fmt.Errorf("folder creation failure: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO dossier_historique (dossier_id, ccial_id, date_start)
		VALUES ($1, $2, $3)
	`, folderId, folder.Ccial_Id, time.Now())

	if err != nil {
		return 0, fmt.Errorf("failed historical addition : %v", err)
	}

	return folderId, nil
}

func GetFolderById(db *sql.DB, id int) (model.Folder, error) {
	var folder model.Folder

	err := db.QueryRow("SELECT * FROM dossier WHERE id = $1", id).Scan(&folder.Id, &folder.Date_insert, &folder.Ccial_Id)

	if err != nil {
		fmt.Println("error")
		return model.Folder{}, err
	}

	return folder, nil
}

func GetAllFolders(db *sql.DB) ([]model.Folder, error) {
	rows, err := db.Query("SELECT id, date_insert, ccial_id FROM dossier")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var folders []model.Folder

	for rows.Next() {
		var f model.Folder
		if err := rows.Scan(&f.Id, &f.Date_insert, &f.Ccial_Id); err != nil {
			return nil, err
		}
		folders = append(folders, f)
	}

	return folders, rows.Err()
}

func UpdateFolderCcial(db *sql.DB, folderId int, ccialId int) error {
	_, err := db.Exec("UPDATE dossier SET ccial_id = $1 WHERE id = $2", ccialId, folderId)
	return err
}

func DeleteFolder(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM dossier WHERE id = $1", id)

	return err
}

func SearchFolders(db *sql.DB, query string) ([]model.Folder, error) {
	rows, err := db.Query("SELECT id, date_insert, ccial_id FROM dossier WHERE CAST(id AS TEXT) LIKE $1 LIMIT 10", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var folders []model.Folder

	for rows.Next() {
		var folder model.Folder
		if err := rows.Scan(&folder.Id, &folder.Date_insert, &folder.Ccial_Id); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}

	return folders, nil
}
