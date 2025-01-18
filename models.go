package PostsModels

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Owner          string    `json:"owner"`
	DateOfCreation time.Time `json:"date_of_creation"`
}
