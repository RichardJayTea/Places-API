package user

import "database/sql"

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

func (u *User) CreateUser(db *sql.DB) error {
	result, err := db.Exec(
		"INSERT INTO teago.user (username, password, email, first_name, last_name, is_active) VALUES (?, ?, ?, ?, ?, ?)",
		u.Username, "somepassword", u.Email, u.FirstName, u.LastName, 1)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()

	return u.GetUserByID(db, int(id))
}

func (u *User) GetUserByID(db *sql.DB, id int) error {
	return db.QueryRow(
		"SELECT uid, username, email, first_name, last_name, is_active, image_path, create_date FROM user WHERE uid = ?",
		id).Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.IsActive, &u.ImagePath, &u.CreateDate)
}
