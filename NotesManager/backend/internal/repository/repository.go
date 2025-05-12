package repository

import (
	"database/sql"
	"time"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type NotesList interface {
	Create(userId int, list domain.NotesList) (int, time.Time, error)
	GetAll(userId int, filter string, sortOrder string) ([]domain.NotesList, error)
	GetById(userId, listId int) (domain.NotesList, error)
	Update(userId, listId int, in domain.UpdateNotesList) (domain.NotesList, error)
	Delete(userId, listId int) (int, error)
}

type NoteItem interface {
	Create(userId, listId int, in domain.UpdateNoteItem) (domain.NoteItem, error)
	GetAll(userId, listId int, filter, sortOrder string) ([]domain.NoteItem, error)
	GetById(userId, itemId int) (domain.NoteItem, error)
	Update(userId, itemId int, in domain.UpdateNoteItem) (domain.NoteItem, error)
	Delete(userId, itemId int) (int, error)
}

type Repository struct {
	Authorization
	NotesList
	NoteItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		NotesList:     NewNotesListPostgres(db),
		NoteItem:      NewNotesPostgres(db),
	}
}
