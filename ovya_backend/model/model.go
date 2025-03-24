package model

import (
	"time"
)

type Acq struct {
	ID       int
	Nom      string
	Email    string
	Password string
}

type Ccial struct {
	Id    int
	Nom   string
	Email string
}

type Folder struct {
	Id          int
	Date_insert time.Time
	Ccial_Id    int
}

type Visite struct {
	Id        int       `json:"id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	AcqId     int       `json:"acq_id"`
	CcialId   int       `json:"ccial_id"`
	DossierId int       `json:"dossier_id"`
	Canceled  bool      `json:"canceled"`
}

type VisiteGroup struct {
	Date       string `json:"date"`
	CcialID    int    `json:"ccial_id"`
	VisiteNb   int    `json:"visite_nb"`
	DossierIds []int  `json:"dossier_ids"`
}

type FoldersHistory struct {
	ID         int        `json:"id"`
	DossierID  int        `json:"dossier_id"`
	CcialID    int        `json:"ccial_id"`
	DateStart  time.Time  `json:"date_start"`
	DateEnd    *time.Time `json:"date_end,omitempty"`
}