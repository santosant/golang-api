package models

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description bool   `json:"description"`
	Completed   bool   `json:"completed"`
}
