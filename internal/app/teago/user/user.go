package user

import (
	"database/sql"
	"fmt"
	"github.com/richardjaytea/teago/cmd/teago/app/requests"
	"github.com/richardjaytea/teago/internal/app/teago/dbhelper"
)

const table = "teago.user"

type User struct {
	ID         int    `json:"uid"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsActive   int    `json:"is_active"`
	ImagePath  string `json:"image_path"`
	CreateDate string `json:"create_date"`
}

func CreateUser(db *sql.DB, data requests.CreateUser) (*User, error) {
	result, err := db.Exec(
		"INSERT INTO teago.user (username, password, email, first_name, last_name, is_active) VALUES (?, ?, ?, ?, ?, ?)",
		data.Username, "somepassword", data.Email, data.FirstName, data.LastName, 1)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return GetUserByID(db, int(id))
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var u User
	err := db.QueryRow(
		"SELECT uid, username, email, first_name, last_name, is_active, image_path, create_date FROM teago.user WHERE uid = ?",
		id).Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.IsActive, &u.ImagePath, &u.CreateDate)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UpdateUserByID(db *sql.DB, id int, data map[string]interface{}) (*User, error) {
	w := fmt.Sprintf("WHERE uid = %d", id)
	s, p := dbhelper.BuildUpdateStatement(table, data, w)

	_, err := db.Exec(s, p...)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return GetUserByID(db, id)
}
