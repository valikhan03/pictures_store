package entity

type User struct {
	UserID   string `json:"user_id" db:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
