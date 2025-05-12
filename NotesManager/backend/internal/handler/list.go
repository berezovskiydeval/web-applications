package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary      Create list
// @Tags         notes-lists
// @Description  создать новый список заметок
// @ID           create-list
// @Accept       json
// @Produce      json
// @Param        input  body      domain.NotesList  true  "данные списка"
// @Success      200    {object}  map[string]int    "list_id"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Входящая модель без CreatedAt (оно заполнится в БД)
	var input domain.NotesList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// теперь сервис отдаёт id и время создания
	idList, createdAt, err := h.services.NotesList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// возвращаем оба поля
	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id":    idList,
		"created_at": createdAt, // время в формате RFC3339
	})
}

type getAllListsResponse struct {
	Data []domain.NotesList `json:"data"`
}

// @Summary      Get all lists
// @Tags         notes-lists
// @Description  получить все списки текущего пользователя
// @ID           get-all-lists
// @Produce      json
// @Success      200  {object}  getAllListsResponse
// @Failure      500  {object}  Error
// @Router       /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	filter := c.DefaultQuery("q", "")          
	sortOrder := c.DefaultQuery("sort", "desc")
	if sortOrder != "asc" && sortOrder != "desc" {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort parameter; must be 'asc' or 'desc'")
		return
	}
	lists, err := h.services.NotesList.GetAll(userID, filter, sortOrder)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary      Get list by ID
// @Tags         notes-lists
// @Description  получить один список по ID
// @ID           get-list-by-id
// @Produce      json
// @Param        id   path      int  true  "list ID"
// @Success      200  {object}  domain.NotesList
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists/{id} [get]
func (h *Handler) getById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.NotesList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary      Update list
// @Tags         notes-lists
// @Description  изменить заголовок / описание списка
// @ID           update-list
// @Accept       json
// @Produce      json
// @Param        id     path      int                    true  "list ID"
// @Param        input  body      domain.UpdateNotesList true  "новые данные"
// @Success      200  {object}  map[string]int  "is_changed (1 — изменено, 0 — нет)"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, _ := strconv.Atoi(c.Param("listId"))

	var input domain.UpdateNotesList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.services.NotesList.Update(userId, listId, input)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusNotFound, "list not found")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, updated)
}

// @Summary      Delete list
// @Tags         notes-lists
// @Description  удалить список (каскадно удалит его заметки)
// @ID           delete-list
// @Produce      json
// @Param        id   path      int  true  "list ID"
// @Success      200  {object}  map[string]int  "is_deleted (1 — удалено, 0 — нет)"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	is_deleted, err := h.services.NotesList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"is_deleted": is_deleted,
	})
}
