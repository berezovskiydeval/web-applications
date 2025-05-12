package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	db "github.com/berezovskyivalerii/notes-manager/backend/pkg/database"
)

type NotesPostgres struct {
	db *sql.DB
}

func NewNotesPostgres(db *sql.DB) *NotesPostgres {
	return &NotesPostgres{
		db: db,
	}
}

func (r *NotesPostgres) Create(userId, listId int, in domain.UpdateNoteItem) (domain.NoteItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SQL с RETURNING для получения id и created_at
	query := fmt.Sprintf(`
        INSERT INTO %s (title, content, pinned, list_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, title, content, pinned, created_at, list_id
    `, db.NotesTable)

	var item domain.NoteItem
	err := r.db.QueryRowContext(
		ctx,
		query,
		in.Title,
		in.Content,
		in.Pinned,
		listId,
	).Scan(
		&item.Id,
		&item.Title,
		&item.Content,
		&item.Pinned,
		&item.CreatedAt,
		&item.ListId,
	)
	if err != nil {
		return domain.NoteItem{}, err
	}
	return item, nil
}

func (r *NotesPostgres) GetAll(userId, listId int, filter, sortOrder string) ([]domain.NoteItem, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    so := strings.ToLower(sortOrder)
    if so != "asc" && so != "desc" {
        return nil, fmt.Errorf("invalid sortOrder %q", sortOrder)
    }

    like := "%" + filter + "%"

    query := fmt.Sprintf(`
        SELECT
            n.id,
            n.title,
            n.content,
            n.pinned,
            n.created_at,
            n.list_id
        FROM %s AS n
        JOIN %s AS l ON l.id = n.list_id
        WHERE
            l.user_id = $1
            AND n.list_id = $2
            AND (n.title   ILIKE $3
              OR  n.content ILIKE $3)
        ORDER BY n.pinned DESC, n.created_at %s
    `, db.NotesTable, db.NotesListsTable, so)

    rows, err := r.db.QueryContext(ctx, query, userId, listId, like)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notes []domain.NoteItem
    for rows.Next() {
        var n domain.NoteItem
        if err := rows.Scan(
            &n.Id,
            &n.Title,
            &n.Content,
            &n.Pinned,
            &n.CreatedAt,
            &n.ListId,
        ); err != nil {
            return nil, err
        }
        notes = append(notes, n)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return notes, nil
}

func (r *NotesPostgres) GetById(userId, itemId int) (domain.NoteItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var note domain.NoteItem

	query := fmt.Sprintf(`
		SELECT  n.id,
		        n.title,
		        n.content,
		        n.pinned,
				n.created_at,
		        n.list_id
		FROM    %s AS n
		JOIN    %s AS l ON l.id = n.list_id
		WHERE   n.id      = $1             
		  AND   l.user_id = $2             
		LIMIT 1;
	`, db.NotesTable, db.NotesListsTable)

	err := r.db.QueryRowContext(ctx, query, itemId, userId).Scan(
		&note.Id,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.ListId,
	)

	if err != nil {
		return domain.NoteItem{}, err
	}

	return note, nil
}

func (r *NotesPostgres) Update(userId, itemId int, in domain.UpdateNoteItem) (domain.NoteItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Обновляем только поля, разрешённые в in, и сразу возвращаем все столбцы
	query := fmt.Sprintf(`
        UPDATE %s AS n
        SET
            title   = $3,
            content = $4,
            pinned  = $5
        FROM %s AS l
        WHERE
            n.id      = $1
            AND n.list_id = l.id
            AND l.user_id = $2
        RETURNING
            n.id,
            n.title,
            n.content,
            n.pinned,
            n.created_at,
            n.list_id
    `, db.NotesTable, db.NotesListsTable)

	var updated domain.NoteItem
	row := r.db.QueryRowContext(ctx, query,
		itemId,
		userId,
		in.Title,
		in.Content,
		in.Pinned,
	)
	if err := row.Scan(
		&updated.Id,
		&updated.Title,
		&updated.Content,
		&updated.Pinned,
		&updated.CreatedAt,
		&updated.ListId,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.NoteItem{}, sql.ErrNoRows
		}
		return domain.NoteItem{}, err
	}

	return updated, nil
}

func (r *NotesPostgres) Delete(userId, itemId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`
		DELETE FROM %s AS n
		USING %s AS l
		WHERE n.id      = $1
		  AND n.list_id = l.id
		  AND l.user_id = $2;
	`, db.NotesTable, db.NotesListsTable)

	res, err := r.db.ExecContext(ctx, query, itemId, userId)
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
