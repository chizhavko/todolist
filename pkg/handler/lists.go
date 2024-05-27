package handler

import (
	"net/http"
	"strconv"

	"github.com/chizhavko/todolist"
	"github.com/gin-gonic/gin"
)

type getAllListsResponse struct {
	Data []todolist.TodoList `json:"data"`
}

type getListByIdResponse struct {
	List todolist.TodoList `json:"list"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	items, err := h.Services.TodoList.AllItems(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{Data: items})
}

func (h *Handler) createList(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	var input todolist.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.Services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateList(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "bad user id value")
		return
	}

	var update todolist.TodoListUpdate
	if err := ctx.BindJSON(&update); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !update.Validate() {
		newErrorResponse(ctx, http.StatusBadRequest, "request empty")
		return
	}

	if err := h.Services.UpdateList(userId, listId, update); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{Status: "success"})
}

func (h *Handler) deleteList(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "bad user id value")
		return
	}

	err = h.Services.TodoList.DeleteList(userId, listId)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{Status: "success"})
}

func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := getUserId(ctx)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "bad user id value")
		return
	}

	list, err := h.Services.TodoList.GetListById(userId, listId)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getListByIdResponse{List: list})
}
