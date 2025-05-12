package domain

import "time"

type NotesList struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UserId      int       `json:"user_id" db:"user_id"`
}

type UpdateNotesList struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type NoteItem struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Pinned    bool      `json:"pinned" db:"pinned"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ListId    int       `json:"list_id" db:"list_id"`
}

type UpdateNoteItem struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Pinned  bool   `json:"pinned" db:"pinned"`
}
