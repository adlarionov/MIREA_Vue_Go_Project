package dto

type RegisterRequestDto struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Name         string `json:"name"`
	Phone        int    `json:"phone"`
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}
