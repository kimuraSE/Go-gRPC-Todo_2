package model

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}

type TodoResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"` 
}
