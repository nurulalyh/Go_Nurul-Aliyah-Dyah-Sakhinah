package handler

type LoginInput struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
