package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary      Create note
// @Tags         notes-items
// @Description  создать заметку в указанном списке
// @ID           create-note
// @Accept       json
// @Produce      json
// @Param        id     path      int               true  "list ID"
// @Param        input  body      domain.NoteItem   true  "данные заметки"
// @Success      200    {object}  map[string]int    "id"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid listId param")
		return
	}

	var input domain.UpdateNoteItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newItem, err := h.services.NoteItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newItem)
}

type getAllItemsResponse struct {
	Data []domain.NoteItem `json:"data"`
}

// @Summary      Get all notes
// @Tags         notes-items
// @Description  получить все заметки списка с фильтрацией и сортировкой
// @ID           get-all-notes
// @Produce      json
// @Param        id     path      int     true   "List ID"
// @Param        q      query     string  false  "Поиск по заголовку или содержимому"
// @Param        sort   query     string  false  "Порядок сортировки по created_at: asc или desc"  Enums(asc, desc)
// @Success      200    {object}  getAllItemsResponse
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
    userID, err := getUserId(c)
    if err != nil {
        return
    }

    listID, err := strconv.Atoi(c.Param("listId"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid listId")
        return
    }

    filter := c.DefaultQuery("q", "")
    sortOrder := strings.ToLower(c.DefaultQuery("sort", "asc"))
    if sortOrder != "asc" && sortOrder != "desc" {
        newErrorResponse(c, http.StatusBadRequest, "invalid sort param")
        return
    }

    items, err := h.services.NoteItem.GetAll(userID, listID, filter, sortOrder)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": items})
}

// @Summary      Get note by ID
// @Tags         notes-items
// @Description  получить одну заметку по ID
// @ID           get-note-by-id
// @Produce      json
// @Param        id   path      int  true  "note ID"
// @Success      200  {object}  domain.NoteItem
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	item, err := h.services.NoteItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary      Update note
// @Tags         notes-items
// @Description  изменить содержание заметки
// @ID           update-note
// @Accept       json
// @Produce      json
// @Param        id     path      int                  true  "note ID"
// @Param        input  body      domain.UpdateNoteItem true "новые данные"
// @Success      200    {object}  map[string]int       "is_changed (1/0)"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
    userId, _ := getUserId(c)
    itemId, _ := strconv.Atoi(c.Param("itemId"))

    var input domain.UpdateNoteItem
    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    updated, err := h.services.NoteItem.Update(userId, itemId, input)
    if err != nil {
        if err == sql.ErrNoRows {
            newErrorResponse(c, http.StatusNotFound, "item not found")
        } else {
            newErrorResponse(c, http.StatusInternalServerError, err.Error())
        }
        return
    }

    c.JSON(http.StatusOK, updated)
}

// @Summary      Delete note
// @Tags         notes-items
// @Description  удалить заметку
// @ID           delete-note
// @Produce      json
// @Param        id   path      int  true  "note ID"
// @Success      200  {object}  map[string]int  "is_deleted (1/0)"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	is_deleted, err := h.services.NoteItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"is_deleted": is_deleted,
	})
}
