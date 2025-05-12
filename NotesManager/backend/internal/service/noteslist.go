package service

import (
	"time"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/berezovskyivalerii/notes-manager/backend/internal/repository"
)

type NotesListService struct {
	repo repository.NotesList
}

func NewNotesListService(repo repository.NotesList) *NotesListService{
	return &NotesListService{repo: repo}
}

func (s *NotesListService) Create(userId int, list domain.NotesList) (int, time.Time, error){
	return s.repo.Create(userId, list)
}

func (s *NotesListService) GetAll(userId int, filter string, sortOrder string) ([]domain.NotesList, error){
	return s.repo.GetAll(userId, filter, sortOrder)
}

func (s *NotesListService) GetById(userId, listId int) (domain.NotesList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *NotesListService) 	Update(userId, listId int, list domain.UpdateNotesList) (domain.NotesList, error) {
	return s.repo.Update(userId, listId, list)
}

func (s *NotesListService) Delete(userId, listId int) (int, error){
	return s.repo.Delete(userId, listId)
}