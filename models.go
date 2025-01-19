package PostsModels

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Post struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Owner          string    `json:"owner"`
	DateOfCreation time.Time `json:"date_of_creation"`
}

type Notification struct {
	Creator     string      `json:"creator"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Receivers   []uuid.UUID `json:"receivers"`
}
