package place

import (
	"database/sql"
	"fmt"
	"github.com/richardjaytea/teago/internal/app/teago/dbhelper"
)

const table = "teago.place"

type Place struct {
	ID          int    `json:"uid"`
	ExternalID  int    `json:"external_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	IsActive    int    `json:"is_active"`
	CreateDate  string `json:"create_date"`
}

func CreatePlace(db *sql.DB, data map[string]interface{}) (*Place, error) {
	stmt :=
		`INSERT INTO teago.place (external_id, name, description, is_active)
		VALUES (?, ?, ?, 1)`

	result, err := db.Exec(stmt, data["external_id"], data["name"], data["description"])
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return GetPlaceById(db, int(id))
}

func GetPlaceById(db *sql.DB, id int) (*Place, error) {
	var p Place
	query :=
		`SELECT uid, name, description, image_path, is_active
		FROM teago.place
		WHERE uid = ?`

	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.ImagePath, &p.IsActive)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func UpdatePlace(db *sql.DB, id int, data map[string]interface{}) (*Place, error) {
	w := fmt.Sprintf("WHERE uid = %d", id)
	stmt, p := dbhelper.BuildUpdateStatement(table, data, w)

	if _, err := db.Exec(stmt, p...); err != nil {
		return nil, err
	}

	return GetPlaceById(db, id)
}
