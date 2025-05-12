package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	db "github.com/berezovskyivalerii/notes-manager/backend/pkg/database"
)

type NotesListPostgres struct {
	db *sql.DB
}

func NewNotesListPostgres(db *sql.DB) *NotesListPostgres {
	return &NotesListPostgres{db: db}
}

func (r *NotesListPostgres) Create(userId int, list domain.NotesList) (int, time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, time.Time{}, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var (
		listID    int
		createdAt time.Time
	)
	insertListSQL := fmt.Sprintf(`
		INSERT INTO %s (title, description, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
		`, db.NotesListsTable)

	if err = tx.QueryRowContext(ctx, insertListSQL, list.Title, list.Description, userId).Scan(&listID, &createdAt); err != nil {
		return 0, time.Time{}, err
	}

	if err = tx.Commit(); err != nil {
		return 0, time.Time{}, err
	}

	return listID, time.Time{}, nil
}

func (r *NotesListPostgres) GetAll(userId int, filter string, sortOrder string) ([]domain.NotesList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	so := strings.ToLower(sortOrder)
	if so != "asc" && so != "desc" {
		return nil, errors.New("invalid sortOrder, must be 'asc' or 'desc'")
	}

	like := "%" + filter + "%"

	query := fmt.Sprintf(`
        SELECT id, title, description, created_at, user_id
          FROM %s
         WHERE user_id = $1
           AND (title   ILIKE $2
             OR description ILIKE $2)
         ORDER BY created_at %s
    `, db.NotesListsTable, so)

	rows, err := r.db.QueryContext(ctx, query, userId, like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []domain.NotesList
	for rows.Next() {
		var l domain.NotesList
		if err := rows.Scan(
			&l.Id,
			&l.Title,
			&l.Description,
			&l.CreatedAt,
			&l.UserId,
		); err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *NotesListPostgres) GetById(userId, listId int) (domain.NotesList, error) {
	var list domain.NotesList

	query := fmt.Sprintf(`SELECT id, title, description, created_at, user_id
	                      FROM %s
	                      WHERE id = $1 AND user_id = $2`, db.NotesListsTable)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.
		QueryRowContext(ctx, query, listId, userId).
		Scan(&list.Id, &list.Title, &list.Description, &list.CreatedAt, &list.UserId)

	return list, err
}

func (r *NotesListPostgres) Update(
	userId, listId int,
	in domain.UpdateNotesList,
) (domain.NotesList, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if in.Title != "" {
		setParts = append(setParts, fmt.Sprintf("title = $%d", argPos))
		args = append(args, in.Title)
		argPos++
	}
	if in.Description != "" {
		setParts = append(setParts, fmt.Sprintf("description = $%d", argPos))
		args = append(args, in.Description)
		argPos++
	}
	if len(setParts) == 0 {
		return domain.NotesList{}, errors.New("nothing to update")
	}

	args = append(args, listId, userId)

	query := fmt.Sprintf(`
        UPDATE %s
        SET %s
        WHERE id = $%d AND user_id = $%d
        RETURNING id, title, description, created_at, user_id
    `,
		db.NotesListsTable,
		strings.Join(setParts, ", "),
		argPos,
		argPos+1,
	)

	// выполняем запрос и сразу сканируем в struct
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updated domain.NotesList
	row := r.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&updated.Id,
		&updated.Title,
		&updated.Description,
		&updated.CreatedAt,
		&updated.UserId,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.NotesList{}, sql.ErrNoRows
		}
		return domain.NotesList{}, err
	}

	return updated, nil
}

func (r *NotesListPostgres) Delete(userId, listId int) (int, error) {
	query := fmt.Sprintf(`DELETE FROM %s 
	                      WHERE id = $1 AND user_id = $2`, db.NotesListsTable)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query, listId, userId)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, sql.ErrNoRows
	}

	return int(rows), nil
}
