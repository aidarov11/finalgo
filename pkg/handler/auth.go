package handler

import (
	"fmt"
	"net/http"

	"github.com/aidarov11/todo-app"
	"github.com/gin-gonic/gin"
)

// Оброботчики url (Они должны имет указатели на *gin.Context)
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	fmt.Print(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password " binding: "required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	fmt.Print(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
