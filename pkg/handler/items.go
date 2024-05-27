package handler

import (
	"net/http"
	"strconv"

	"github.com/chizhavko/todolist"
	"github.com/gin-gonic/gin"
)

type getAllItemsResponse struct {
	Data []todolist.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "bad list id param")
		return
	}

	items, err := h.Services.TodoItem.AllItems(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllItemsResponse{Data: items})
}

func (h *Handler) createItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "bad list id param")
		return
	}

	var input todolist.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.TodoItem.Create(userId, listId, input)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateItem(ctx *gin.Context) {

}

func (h *Handler) deleteItem(ctx *gin.Context) {

}

func (h *Handler) getItemById(ctx *gin.Context) {

}
