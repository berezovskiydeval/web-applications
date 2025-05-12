package service

import (
	"time"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/berezovskyivalerii/notes-manager/backend/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	NotesList
	NoteItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		NotesList:     NewNotesListService(repos.NotesList),
		NoteItem:      NewNotesService(repos.NoteItem, repos.NotesList),
	}
}
