package dto

type UserRegisterDto struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	FullName     string `json:"name"`
	Role         string `json:"role"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	Token string `json:"token"`
}
