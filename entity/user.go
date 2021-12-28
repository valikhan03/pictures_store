package entity

type User struct {
	UserID   string `json:"user_id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
