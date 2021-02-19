package requests

// TODO: create middleware to validate request body data against db tag

type CreateUser struct {
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type UpdateUser struct {
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	IsActive  int    `json:"is_active" db:"is_active"`
	ImagePath string `json:"image_path" db:"image_path"`
}
