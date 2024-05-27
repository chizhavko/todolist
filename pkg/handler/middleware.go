package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdContext       = "userId"
)

func (h *Handler) identifyUser(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty header")
		return
	}

	headerPaths := strings.Split(header, " ")

	if len(headerPaths) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "incorrect authorization data in header")
		return
	}

	userId, err := h.Services.Authorization.ParseToken(headerPaths[1])

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userIdContext, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userIdContext)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	return id.(int), nil
}
