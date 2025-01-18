package models

type Post struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Owner          string `json:"owner"`
	DateOfCreation int64  `json:"date_of_creation"`
}
