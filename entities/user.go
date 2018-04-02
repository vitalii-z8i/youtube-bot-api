package entities

// User contains basic info about message sender
type User struct {
	ID        int64  `json:"id" db:"ID"`
	IsBot     bool   `json:"is_bot" db:"-"`
	FirstName string `json:"first_name" db:"FirstName"`
}
