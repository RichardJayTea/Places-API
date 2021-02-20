package checkin

import "database/sql"

const table = "teago.checkin"

type Checkin struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	PlaceID    int    `json:"place_id"`
	CreateDate string `json:"create_date"`
}

func CreateCheckin(db *sql.DB, data map[string]interface{}) error {
	stmt := "INSERT INTO teago.checkin (user_id, place_id) VALUES (?, ?)"
	_, err := db.Exec(stmt, data["user_id"], data["place_id"])

	return err
}

func GetAllCheckinByID(db *sql.DB, id int) ([]interface{}, error) {
	type result struct {
		ID          int    `json:"id"`
		PlaceID     int    `json:"place_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImagePath   string `json:"image_path"`
		CreateDate  string `json:"create_date"`
	}

	stmt :=
		`SELECT c.uid, c.place_id, p.name, p.image_path, c.create_date
		FROM teago.checkin c
		JOIN teago.place p
			ON c.place_id = p.uid
		WHERE c.user_id = ?`
	rows, err := db.Query(stmt, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var rl []interface{}
	var r result

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&r.ID, &r.PlaceID, &r.Name, &r.ImagePath, &r.CreateDate)
		rl = append(rl, r)
	}

	return rl, nil
}
