package request

type AuthLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
