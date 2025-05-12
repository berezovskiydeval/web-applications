package handler

import (
	"net/http"

	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary      SignUp
// @Tags         auth
// @Description  create account
// @ID           create-account
// @Accept       json
// @Produce      json
// @Param        input body domain.User true "account info"
// @Success      200 {integer} integer 1
// @Failure      400,404 {object} Error
// @Failure      500 {object} Error
// @Failure      default {object} Error
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary      SignIn
// @Tags         auth
// @Description  авторизация: возвращает JWT‑токен
// @ID           sign-in
// @Accept       json
// @Produce      json
// @Param        input  body      domain.UserSignIn  true  "учётные данные"
// @Success      200    {string}  string             "accessToken"
// @Failure      400,404  {object}  Error
// @Failure      500       {object}  Error
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input domain.UserSignIn

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token,
	})
}
