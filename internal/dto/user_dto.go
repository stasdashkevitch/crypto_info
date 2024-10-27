package dto

type RegisterUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
