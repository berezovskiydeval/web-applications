package service

import (
	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/berezovskyivalerii/notes-manager/backend/internal/repository"
)

type NotesService struct {
	itemRepo repository.NoteItem
	listRepo repository.NotesList
}

func NewNotesService(itemRepo repository.NoteItem, listRepo repository.NotesList) *NotesService {
	return &NotesService{
		itemRepo: itemRepo,
		listRepo: listRepo,
	}
}

func (s *NotesService) 	Create(userId, listId int, in domain.UpdateNoteItem) (domain.NoteItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return domain.NoteItem{}, err
	}

	return s.itemRepo.Create(userId, listId, in)
}

func (s *NotesService) GetAll(userId, listId int, filter, sortOrder string) ([]domain.NoteItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}

	return s.itemRepo.GetAll(userId, listId, filter, sortOrder)
}

func (s *NotesService) GetById(userId, itemId int) (domain.NoteItem, error) {
	return s.itemRepo.GetById(userId, itemId)
}

func (s *NotesService) Update(userId, itemId int, in domain.UpdateNoteItem) (domain.NoteItem, error){
	return s.itemRepo.Update(userId, itemId, in)
}

func (s *NotesService) Delete(userId, itemId int) (int, error){
	return s.itemRepo.Delete(userId, itemId)
}

