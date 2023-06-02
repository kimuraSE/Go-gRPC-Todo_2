package model

type TodoRequest struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}

type TodoResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	UserId uint   `json:"user_id"` 
}
