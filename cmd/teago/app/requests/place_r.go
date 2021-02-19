package requests

type CreatePlace struct {
	ExternalID  int    `json:"external_id" db:"external_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
