package model

type UserRequest struct {
	ID	   int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Token string `json:"token"`
}
