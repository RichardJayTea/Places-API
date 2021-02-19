package requests

type CreateUser struct {
	Username  string `json:"username" dbhelper:"username"`
	Email     string `json:"email" dbhelper:"email"`
	FirstName string `json:"first_name" dbhelper:"first_name"`
	LastName  string `json:"last_name" dbhelper:"last_name"`
	IsActive  int    `json:"is_active" dbhelper:"is_active"`
	ImagePath string `json:"image_path" dbhelper:"image_path"`
}

type UpdateUser struct {
	Username  string `json:"username" dbhelper:"username"`
	Email     string `json:"email" dbhelper:"email"`
	FirstName string `json:"first_name" dbhelper:"first_name"`
	LastName  string `json:"last_name" dbhelper:"last_name"`
	IsActive  int    `json:"is_active" dbhelper:"is_active"`
	ImagePath string `json:"image_path" dbhelper:"image_path"`
}
