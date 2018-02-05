package entities

// User contains basic info about message sender
type User struct {
	ID        int32  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
}
