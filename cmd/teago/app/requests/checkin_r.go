package requests

type CreateCheckin struct {
	UserID  int `json:"user_id" db:"user_id"`
	PlaceID int `json:"place_id" db:"user_id"`
}
