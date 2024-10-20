package dtos

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
